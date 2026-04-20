package rest_err

import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message"`
	Code    int      `json:"code"`
	Summary string   `json:"resume"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message, sumarry string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Summary: sumarry,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRedquestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Summary: "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Summary: "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewBadInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Summary: "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Summary: "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Summary: "Forbidden",
		Code:    http.StatusForbidden,
	}
}
