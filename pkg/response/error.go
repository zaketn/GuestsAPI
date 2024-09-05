package response

import (
	"fmt"
	"log"
	"net/http"
)

func BadRequest() []byte {
	return Make(WithCode(http.StatusBadRequest))
}

func InternalServerError() []byte {
	return Make(WithCode(http.StatusInternalServerError))
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

	fmt.Fprintln(w, error)
}
