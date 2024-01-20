package models

import "github.com/go-raptor/raptor"

func NewError(code int, message string) raptor.Map {
	return raptor.Map{
		"code":    code,
		"message": message,
	}
}
