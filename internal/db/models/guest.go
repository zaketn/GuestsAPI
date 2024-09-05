package models

import (
	"github.com/jackc/pgx"
)

type Guest struct {
	Id                                    int
	Name, LastName, Email, Phone, Country string
}

type GuestModel struct {
	DB *pgx.Conn
}

func (m *GuestModel) Create(guest *Guest) (*Guest, error) {
	sql := `INSERT INTO guests (name, last_name, email, phone, country) VALUES ($1, $2, $3, $4, $5) RETURNING id`

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
	sql := "SELECT * FROM guests"
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
	sql := `UPDATE guests SET 
                  name = $1,
                  last_name = $2,
                  email = $3,
                  phone = $4,
                  country = $5
              WHERE id = $6
              RETURNING id, name, last_name, email, phone, country`

	updatedGuest := &Guest{}

	err := m.DB.QueryRow(sql,
		newData.Name,
		newData.LastName,
		newData.Email,
		newData.Phone,
		newData.Country,
		newData.Id).Scan(
		&updatedGuest.Id,
		&updatedGuest.Name,
		&updatedGuest.LastName,
		&updatedGuest.Email,
		&updatedGuest.Phone,
		&updatedGuest.Country)

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
