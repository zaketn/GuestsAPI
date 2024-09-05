package response

import "net/http"

type RespOpts struct {
	Status  string                 `json:"status,omitempty"`
	Code    int                    `json:"code,omitempty"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   *Response              `json:"error,omitempty"`
	Details []map[string]string    `json:"details,omitempty"`
}

type OptsFunc func(options *RespOpts)

func defaultOptions() RespOpts {
	return RespOpts{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
	}
}

func WithCode(code int) OptsFunc {
	return func(options *RespOpts) {
		options.Code = code
		options.Status = http.StatusText(code)
	}
}

func WithMessage(message string) OptsFunc {
	return func(options *RespOpts) {
		options.Message = message
	}
}

func WithData(data interface{}) OptsFunc {
	return func(options *RespOpts) {
		dataMap := make(map[string]interface{})
		dataMap["result"] = data
		options.Data = dataMap
	}
}

func WithNamedData(dataName string, data interface{}) OptsFunc {
	return func(options *RespOpts) {
		dataMap := make(map[string]interface{})
		dataMap[dataName] = data
		options.Data = dataMap
	}
}

func WithDetails(details []map[string]string) OptsFunc {
	return func(options *RespOpts) {
		options.Details = details
	}
}
