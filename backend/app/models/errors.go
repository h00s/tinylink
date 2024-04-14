package models

import (
	"net/http"

	"github.com/go-raptor/raptor"
)

func NewError(code int, message string) raptor.Map {
	return raptor.Map{
		"code":    code,
		"message": message,
	}
}

func NewErrorNotFound() (raptor.Map, int) {
	return NewError(404, "Not found"), http.StatusNotFound
}

func NewErrorUnauthorized() (raptor.Map, int) {
	return NewError(401, "Unauthorized"), http.StatusUnauthorized
}

func NewErrorInternal() (raptor.Map, int) {
	return NewError(500, "Internal error"), http.StatusInternalServerError
}

func NewErrorBadRequest() (raptor.Map, int) {
	return NewError(400, "Bad request"), http.StatusBadRequest
}
