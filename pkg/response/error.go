package response

import (
	"log"
	"net/http"
)

func ReturnBadRequestError(w http.ResponseWriter, err error) {
	log.Println(err)

	res := Make(WithCode(http.StatusBadRequest), WithMessage(err.Error()))
	jsonError(w, res, http.StatusBadRequest)
}

func ReturnNotFoundError(w http.ResponseWriter, err error) {
	log.Println(err)

	res := Make(WithCode(http.StatusNotFound), WithMessage(err.Error()))
	jsonError(w, res, http.StatusNotFound)
}

func ReturnInternalError(w http.ResponseWriter, err error) {
	log.Println(err)

	res := Make(WithCode(http.StatusInternalServerError))
	jsonError(w, res, http.StatusInternalServerError)
}

func ReturnError(w http.ResponseWriter, resContent []byte, code int) {
	h := w.Header()
	h.Set("Content-Type", "application/json")

	jsonError(w, resContent, code)
}

func jsonError(w http.ResponseWriter, error []byte, code int) {
	h := w.Header()

	h.Del("Content-Length")
	h.Set("Content-Type", "application/json")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	w.Write(error)
}
