package main

import (
	"context"
	"github.com/jackc/pgx"
	"github.com/zaketn/GuestsAPI/internal/db/models"
	"log"
	"os"
	"strconv"
)

type application struct {
	db    *pgx.Conn
	guest *models.GuestModel
}

func main() {
	conn, err := dbConnect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	app := &application{
		db:    conn,
		guest: &models.GuestModel{DB: conn},
	}

	guests, err := app.guest.GetAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, guest := range guests {
		log.Println(guest)
	}
}

func dbConnect() (*pgx.Conn, error) {
	dbPort, err := strconv.ParseUint(os.Getenv("PG_PORT"), 10, 32)
	if err != nil {
		return nil, err
	}

	dbConfig := pgx.ConnConfig{
		Host:     "db",
		Port:     uint16(dbPort),
		Database: os.Getenv("PG_DB_NAME"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
	}

	conn, err := pgx.Connect(dbConfig)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
