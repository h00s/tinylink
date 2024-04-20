package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
)

type AccessesService struct {
	raptor.Service
}

func (as *AccessesService) Create(linkID uint) error {
	if linkID > 0 {
		return as.DB.Create(&models.Access{LinkID: linkID}).Error
	}
	return raptor.NewErrorBadRequest("Invalid link ID")
}
