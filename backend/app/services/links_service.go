package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/internal"
)

type LinksService struct {
	raptor.Service
}

func (l *LinksService) Get(id uint) (models.Link, error) {
	var link models.Link
	if err := l.DB.First(&link, id).Error; err != nil {
		return link, err
	}
	return link, nil
}

func (l *LinksService) GetByShortID(shortID string) (models.Link, error) {
	return l.Get(internal.IDFromShortURI(shortID))
}

func (l *LinksService) GetByURL(url string) (models.Link, error) {
	var link models.Link
	if err := l.DB.
		Where("url = ?", url).
		First(&link).Error; err != nil {
		return link, err
	}
	return link, nil
}

func (l *LinksService) Create(link models.Link) (models.Link, error) {
	if err := internal.IsURLValid(link.URL); err != nil {
		return link, err
	}
	// find if url already exists
	err := l.DB.
		Select(models.LinkPermittedParams).
		Create(&link).Error
	if err != nil {
		return link, err
	}
	return l.Get(link.ID)
}
