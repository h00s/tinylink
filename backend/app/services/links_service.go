package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/internal"
)

type LinksService struct {
	raptor.Service
}

func (ls *LinksService) Get(id uint) (models.Link, error) {
	var link models.Link
	if err := ls.DB.First(&link, id).Error; err != nil {
		return link, err
	}
	return link, nil
}

func (ls *LinksService) GetByShortID(shortID string) (models.Link, error) {
	return ls.Get(internal.IDFromShortURI(shortID))
}

func (ls *LinksService) GetByURL(url string) (models.Link, error) {
	var link models.Link
	if err := ls.DB.
		Where("url = ?", url).
		First(&link).Error; err != nil {
		return link, err
	}
	return link, nil
}

func (ls *LinksService) Create(link models.Link) (models.Link, error) {
	if err := internal.IsURLValid(link.URL); err != nil {
		return link, err
	}
	if l, err := ls.GetByURL(link.URL); err == nil {
		return l, nil
	}
	err := ls.DB.
		Select(models.LinkPermittedParams).
		Create(&link).Error
	if err != nil {
		return link, err
	}
	return ls.Get(link.ID)
}
