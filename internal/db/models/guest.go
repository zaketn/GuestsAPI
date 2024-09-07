package models

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"strings"
)

type Guest struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Country  string `json:"country"`
}

type GuestModel struct {
	DB *pgx.Conn
}

func (m *GuestModel) Create(guest *Guest) (*Guest, error) {
	sql := `INSERT INTO guests (name, last_name, email, phone, country)
			VALUES (INITCAP($1), INITCAP($2), $3, $4, $5) RETURNING id`

	err := m.DB.QueryRow(sql, guest.Name, guest.LastName, guest.Email, guest.Phone, guest.Country).Scan(&guest.Id)

	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (m *GuestModel) Get(id int) (*Guest, error) {
	sql := "SELECT * FROM guests WHERE id = $1"
	guest := &Guest{}

	err := m.DB.QueryRow(sql, id).Scan(
		&guest.Id,
		&guest.Name,
		&guest.LastName,
		&guest.Email,
		&guest.Phone,
		&guest.Country)

	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (m *GuestModel) GetAll() ([]*Guest, error) {
	sql := "SELECT * FROM guests ORDER BY id"
	guests := make([]*Guest, 0)

	query, err := m.DB.Query(sql)

	if err != nil {
		return nil, err
	}

	for query.Next() {
		guest := &Guest{}

		err := query.Scan(&guest.Id, &guest.Name, &guest.LastName, &guest.Email, &guest.Phone, &guest.Country)

		if err != nil {
			return nil, err
		}

		guests = append(guests, guest)
	}

	return guests, nil
}

func (m *GuestModel) Update(newData *Guest) (*Guest, error) {
	fields := []struct {
		name     string
		value    string
		modifier string
	}{
		{"name", newData.Name, "INITCAP"},
		{"last_name", newData.LastName, "INITCAP"},
		{"email", newData.Email, ""},
		{"phone", newData.Phone, ""},
		{"country", newData.Country, "UPPER"},
	}

	sql := `UPDATE guests SET`
	params := []interface{}{}
	paramIdx := 1

	for _, field := range fields {
		if field.value != "" {
			if field.modifier != "" {
				sql += fmt.Sprintf(" %s = %s($%d),", field.name, field.modifier, paramIdx)
			} else {
				sql += fmt.Sprintf(" %s = $%d,", field.name, paramIdx)
			}
			params = append(params, field.value)
			paramIdx++
		}
	}

	if len(params) == 0 {
		return nil, errors.New("no fields to update")
	}

	sql = strings.TrimSuffix(sql, ",")
	sql += fmt.Sprintf(" WHERE id = $%d RETURNING id, name, last_name, email, phone, country", paramIdx)
	params = append(params, newData.Id)

	updatedGuest := &Guest{}

	err := m.DB.QueryRow(sql, params...).Scan(
		&updatedGuest.Id,
		&updatedGuest.Name,
		&updatedGuest.LastName,
		&updatedGuest.Email,
		&updatedGuest.Phone,
		&updatedGuest.Country,
	)

	if err != nil {
		return nil, err
	}

	return updatedGuest, nil
}

func (m *GuestModel) Delete(id int) (int, error) {
	sql := "DELETE FROM guests WHERE id = $1 RETURNING id"

	deletedGuestId := 0

	err := m.DB.QueryRow(sql, id).Scan(&deletedGuestId)

	if err != nil {
		return 0, err
	}

	return deletedGuestId, nil
}
