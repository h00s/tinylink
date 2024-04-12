package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/internal"
)

type LinksService struct {
	raptor.Service
}

func (l *LinksService) Get(id uint) (models.PublicLink, error) {
	var link models.Link
	if err := l.DB.First(&link, id).Error; err != nil {
		return link.ToPublicLink(), err
	}
	return link.ToPublicLink(), nil
}

func (l *LinksService) GetByShortID(shortID string) (models.PublicLink, error) {
	return l.Get(internal.IDFromShortURI(shortID))
}

func (l *LinksService) Create(link models.Link) (models.PublicLink, error) {
	err := l.DB.
		Select(models.LinkPermittedParams).
		Create(&link).Error
	if err != nil {
		return link.ToPublicLink(), err
	}
	return l.Get(link.ID)
}
