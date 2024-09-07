package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zaketn/GuestsAPI/pkg/response"
	"net/http"
)

func router(app *application) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Get("/guests", app.index)

	r.Post("/guest", app.createGuest)
	r.Get("/guest/", app.getGuest)
	r.Patch("/guest", app.updateGuest)
	r.Delete("/guest", app.deleteGuest)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		response.ReturnError(w, response.Make(response.WithCode(http.StatusNotFound)), http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		response.ReturnError(
			w,
			response.Make(response.WithCode(http.StatusMethodNotAllowed)),
			http.StatusMethodNotAllowed)
	})

	return r
}
