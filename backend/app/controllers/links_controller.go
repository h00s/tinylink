package controllers

import (
	"errors"
	"net/http"

	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/app/services"
	"gorm.io/gorm"
)

type LinksController struct {
	raptor.Controller

	Links *services.LinksService
}

func (lc *LinksController) Get(c *raptor.Context) error {
	link, err := lc.Links.GetByShortID(c.Params("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(models.NewErrorNotFound())
		} else {
			return c.JSON(models.NewErrorInternal())
		}
	}
	if link.Password != "" {
		return c.JSON(models.NewErrorUnauthorized())
	}
	return c.JSON(link.ToPublicLink())
}

func (lc *LinksController) Create(c *raptor.Context) error {
	var link models.Link
	if err := c.BodyParser(&link); err != nil {
		return c.JSON(models.NewErrorBadRequest())
	}
	l, err := lc.Links.Create(link)
	if err != nil {
		return c.JSON(models.NewErrorInternal())
	}
	return c.JSON(l.ToPublicLink(), http.StatusCreated)
}
