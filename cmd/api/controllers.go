package main

import (
	"github.com/zaketn/GuestsAPI/internal/db/models"
	"github.com/zaketn/GuestsAPI/pkg/response"
	"github.com/zaketn/GuestsAPI/pkg/validation"
	"net/http"
	"strconv"
)

func (app application) index(w http.ResponseWriter, r *http.Request) {
	guests, err := app.guest.GetAll()
	if err != nil {
		response.ReturnInternalError(w, err)
		return
	}

	res := response.Make(
		response.WithNamedData("guests", guests),
		response.WithMessage("The list of all users."),
	)

	w.Write(res)
}

func (app application) createGuest(w http.ResponseWriter, r *http.Request) {
	err := validation.FormValidator{Request: r}.Validate(&validation.Ruleset{
		Rules: &map[string][]validation.Rule{
			"name": {
				validation.NotEmpty(),
				validation.Length(1, 128),
				validation.String(),
			},
			"last_name": {
				validation.NotEmpty(),
				validation.Length(1, 128),
			},
			"email": {
				validation.NotEmpty(),
				validation.Email(),
				validation.DoesNotExist(app.db, "guests", "email"),
			},
			"phone": {
				validation.NotEmpty(),
				validation.Phone(),
				validation.DoesNotExist(app.db, "guests", "phone"),
			},
			"country": {
				validation.CountryCode(),
			},
		}})
	if err != nil {
		response.ReturnBadRequestError(w, err)
		return
	}

	newGuest := models.Guest{
		Name:     r.PostForm.Get("name"),
		LastName: r.PostForm.Get("last_name"),
		Email:    r.PostForm.Get("email"),
		Phone:    r.PostForm.Get("phone"),
		Country:  r.PostForm.Get("country"),
	}

	if newGuest.Country == "" {
		country, err := matchCountryFromPhone(newGuest.Phone)
		if err != nil {
			response.ReturnInternalError(w, err)
			return
		}

		newGuest.Country = country
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
	err := validation.FormValidator{Request: r}.Validate(&validation.Ruleset{
		Rules: &map[string][]validation.Rule{
			"id": {
				validation.NotEmpty(),
				validation.Numeric(),
				validation.Exists(app.db, "guests", "id"),
			},
		}})
	if err != nil {
		response.ReturnNotFoundError(w, err)
		return
	}

	guestId, _ := strconv.Atoi(r.Form.Get("id"))
	guest, err := app.guest.Get(guestId)
	if err != nil {
		response.ReturnInternalError(w, err)
		return
	}

	res := response.Make(
		response.WithNamedData("guest", guest),
		response.WithMessage("The user was successfully retrieved."),
	)

	w.Write(res)
}

func (app application) updateGuest(w http.ResponseWriter, r *http.Request) {
	err := validation.FormValidator{Request: r}.Validate(&validation.Ruleset{
		Rules: &map[string][]validation.Rule{
			"id": {
				validation.NotEmpty(),
				validation.Numeric(),
				validation.Exists(app.db, "guests", "id"),
			},
			"name": {
				validation.NotEmpty(),
				validation.Length(1, 128),
				validation.String(),
			},
			"last_name": {
				validation.NotEmpty(),
				validation.Length(1, 128),
			},
			"email": {
				validation.NotEmpty(),
				validation.Email(),
				validation.DoesNotExist(app.db, "guests", "email"),
			},
			"phone": {
				validation.NotEmpty(),
				validation.Phone(),
				validation.DoesNotExist(app.db, "guests", "phone"),
			},
			"country": {
				validation.NotEmpty(),
				validation.CountryCode(),
			},
		}})
	if err != nil {
		response.ReturnBadRequestError(w, err)
		return
	}

	guestId, _ := strconv.Atoi(r.Form.Get("id"))
	updatedGuest := &models.Guest{
		Id:       guestId,
		Name:     r.PostForm.Get("name"),
		LastName: r.PostForm.Get("last_name"),
		Email:    r.PostForm.Get("email"),
		Phone:    r.PostForm.Get("phone"),
		Country:  r.PostForm.Get("country"),
	}

	guest, err := app.guest.Update(updatedGuest)
	if err != nil {
		response.ReturnInternalError(w, err)
		return
	}

	res := response.Make(response.WithNamedData("guest", guest), response.WithMessage("The user was successfully updated."))

	w.Write(res)
}

func (app application) deleteGuest(w http.ResponseWriter, r *http.Request) {
	err := validation.FormValidator{Request: r}.Validate(&validation.Ruleset{
		Rules: &map[string][]validation.Rule{
			"id": {
				validation.NotEmpty(),
				validation.Numeric(),
				validation.Exists(app.db, "guests", "id"),
			},
		}})
	if err != nil {
		response.ReturnNotFoundError(w, err)
		return
	}

	guestId, _ := strconv.Atoi(r.Form.Get("id"))
	guest, err := app.guest.Delete(guestId)
	if err != nil {
		response.ReturnInternalError(w, err)
		return
	}

	res := response.Make(response.WithData(guest), response.WithMessage("The user was successfully deleted."))

	w.Write(res)
}
