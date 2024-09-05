package main

import (
	"github.com/jackc/pgx"
	"github.com/zaketn/GuestsAPI/internal/db/models"
	"log"
	"net/http"
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

	err = http.ListenAndServe(":8000", router(app))

	log.Panic(err)
}
