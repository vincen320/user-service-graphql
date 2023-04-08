package response

import (
	"encoding/json"
	"net/http"
)

type (
	Response struct {
		Code   int `json:"-"`
		Data   any `json:"data,omitempty"`
		Errors any `json:"errors,omitempty"`
	}
)

func NewResponse(code int, res any) (response Response) {
	response.Code = code
	if code >= http.StatusOK && code < 300 {
		response.Data = res
	} else {
		response.Errors = res
	}
	return
}

// JSON sends a JSON response & escape HTML Tags
func (res Response) JSON(w http.ResponseWriter) (err error) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(res.Code)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(true)
	return encoder.Encode(res)
}

// JSONUnsafe sends a JSON response & not escape HTML Tags
func (res Response) JSONUnsafe(w http.ResponseWriter) (err error) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(res.Code)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(res)
}
