package response

import (
	"log"
	"net/http"
)

func ReturnBadRequestError(w http.ResponseWriter, err error) {
	log.Println(err)

	res := Make(WithCode(http.StatusBadRequest), WithMessage(err.Error()))
	jsonError(w, string(res), http.StatusBadRequest)
}

func ReturnNotFoundError(w http.ResponseWriter, err error) {
	log.Println(err)

	res := Make(WithCode(http.StatusNotFound), WithMessage(err.Error()))
	jsonError(w, string(res), http.StatusNotFound)
}

func ReturnInternalError(w http.ResponseWriter, err error) {
	log.Println(err)

	res := Make(WithCode(http.StatusInternalServerError))
	jsonError(w, string(res), http.StatusInternalServerError)
}

func ReturnError(w http.ResponseWriter, resContent []byte, code int) {
	h := w.Header()
	h.Set("Content-Type", "application/json")

	jsonError(w, string(resContent), code)
}

func jsonError(w http.ResponseWriter, error string, code int) {
	h := w.Header()

	h.Del("Content-Length")
	h.Set("Content-Type", "application/json")
	h.Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
}
