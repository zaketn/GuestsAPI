package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/zaketn/GuestsAPI/internal/db/models"
	"github.com/zaketn/GuestsAPI/pkg/response"
	"net/http"
	"strconv"
)

func (app application) index(w http.ResponseWriter, r *http.Request) {
	guests, err := app.guest.GetAll()
	if err != nil {
		response.ReturnError(w, response.BadRequest(), http.StatusBadRequest)
		return
	}

	res := response.Make(response.WithNamedData("guests", guests))

	w.Write(res)
}

func (app application) createGuest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.ReturnError(w, response.BadRequest(), http.StatusBadRequest)
		return
	}

	newGuest := models.Guest{
		Name:     r.Form.Get("name"),
		LastName: r.Form.Get("last_name"),
		Email:    r.Form.Get("email"),
		Phone:    r.Form.Get("phone"),
		Country:  r.Form.Get("country"),
	}

	guest, err := app.guest.Create(&newGuest)
	if err != nil {
		response.ReturnInternalError(w, err)
		return
	}

	res := response.Make(
		response.WithCode(http.StatusCreated),
		response.WithNamedData("guest", guest),
		response.WithMessage("The user was successfully created."),
	)

	w.Write(res)
}

func (app application) getGuest(w http.ResponseWriter, r *http.Request) {
	guestId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ReturnError(w, response.BadRequest(), http.StatusBadRequest)
		return
	}

	guest, err := app.guest.Get(guestId)
	if err != nil {
		response.ReturnError(w, response.Make(response.WithCode(404)), http.StatusNotFound)
		return
	}

	res := response.Make(response.WithNamedData("guest", guest))

	w.Write(res)
}

func (app application) updateGuest(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.ReturnError(w, response.BadRequest(), http.StatusBadRequest)
		return
	}

	guestId, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		response.ReturnError(w, response.BadRequest(), http.StatusBadRequest)
		return
	}

	updatedGuest := &models.Guest{
		Id:       guestId,
		Name:     r.Form.Get("name"),
		LastName: r.Form.Get("last_name"),
		Email:    r.Form.Get("email"),
		Phone:    r.Form.Get("phone"),
		Country:  r.Form.Get("country"),
	}

	guest, err := app.guest.Update(updatedGuest)
	if err != nil {
		response.InternalServerError()
		return
	}

	res := response.Make(response.WithNamedData("guest", guest), response.WithMessage("The user was successfully updated."))

	w.Write(res)
}

func (app application) deleteGuest(w http.ResponseWriter, r *http.Request) {
	guestId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ReturnError(w, response.BadRequest(), http.StatusBadRequest)
		return
	}

	guest, err := app.guest.Delete(guestId)
	if err != nil {
		response.ReturnError(w, response.Make(response.WithCode(404)), http.StatusNotFound)
		return
	}

	res := response.Make(response.WithData(guest))

	w.Write(res)
}
