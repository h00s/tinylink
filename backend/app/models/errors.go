package models

import "github.com/go-raptor/raptor"

func NewError(code int, message string) raptor.Map {
	return raptor.Map{
		"code":    code,
		"message": message,
	}
}

func NewErrorNotFound() raptor.Map {
	return NewError(404, "Not found")
}

func NewErrorUnauthorized() raptor.Map {
	return NewError(401, "Unauthorized")
}

func NewErrorInternal() raptor.Map {
	return NewError(500, "Internal error")
}

func NewErrorBadRequest() raptor.Map {
	return NewError(400, "Bad request")
}
