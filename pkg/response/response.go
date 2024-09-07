package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	RespOpts
}

func Make(responseOptions ...OptsFunc) []byte {
	options := defaultOptions()

	for _, fn := range responseOptions {
		fn(&options)
	}

	res := &Response{options}

	if res.Code >= 400 {
		res = &Response{
			RespOpts{
				Error: res,
			},
		}
	}

	return toJson(res)
}

func toJson(response *Response) []byte {
	jsonRes, err := json.Marshal(response)

	if err != nil {
		log.Println(err)

		errRes := Make(WithCode(http.StatusInternalServerError))
		jsonErrRes, _ := json.Marshal(errRes)

		return jsonErrRes
	}

	return jsonRes
}
