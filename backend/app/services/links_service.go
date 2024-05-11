package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/internal"
	"gorm.io/gorm"
)

type LinksService struct {
	raptor.Service

	Errors   *ErrorsService
	Memstore *Memstore
}

func (ls *LinksService) Get(id uint) (models.Link, error) {
	var link models.Link
	data, err := ls.Memstore.Get(fmt.Sprintf("links:%d", id))
	if err == nil && data != "" {
		json.Unmarshal([]byte(data), &link)
		return link, nil
	}
	if err := ls.DB.
		Where("valid = ?", true).
		First(&link, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return link, raptor.NewErrorNotFound()
		} else {
			return link, raptor.NewErrorInternal()
		}
	}
	go ls.Memstore.Set(fmt.Sprintf("links:%d", id), link)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return link, raptor.NewErrorNotFound()
		} else {
			return link, raptor.NewErrorInternal()
		}
	}
	return link, nil
}

func (ls *LinksService) Create(link models.Link) (models.Link, error) {
	if err := internal.IsURLValid(link.URL); err != nil {
		return link, raptor.NewErrorBadRequest(ls.Errors.Message(err.Error()))
	}
	link.Valid = true

	if l, err := ls.GetByURL(link.URL); err == nil && link.Password == "" {
		if !l.Valid {
			return link, raptor.NewErrorForbidden(ls.Errors.Message("LINK_NOT_VALID"))
		}
		return l, nil
	}

	if link.Password != "" {
		hashedPassword, err := internal.HashPassword(link.Password)
		if err != nil {
			return link, raptor.NewErrorInternal(ls.Errors.Message("ERROR_HASHING_PASSWORD"))
		}
		link.Password = hashedPassword
	}

	err := ls.DB.
		Select(models.LinkPermittedParams).
		Create(&link).Error
	if err != nil {
		return link, raptor.NewErrorInternal(ls.Errors.Message("ERROR_CREATING_LINK"))
	}

	return ls.Get(link.ID)
}
