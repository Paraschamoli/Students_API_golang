package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var messages []string

	for _, err := range errs {
		field := err.Field()
		if err.ActualTag() == "required" {
			messages = append(messages, fmt.Sprintf("field %s is required", field))
		} else {
			messages = append(messages, fmt.Sprintf("field %s is invalid", field))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(messages, ", "),
	}
}
