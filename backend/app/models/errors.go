package models

import "github.com/h00s/raptor"

func NewError(code int, message string) raptor.Map {
	return raptor.Map{
		"code":    code,
		"message": message,
	}
}
