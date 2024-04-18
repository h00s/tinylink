package controllers

import (
	"net/http"

	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/app/services"
)

type LinksController struct {
	raptor.Controller

	Links *services.LinksService
}

func (lc *LinksController) Get(c *raptor.Context) error {
	link, err := lc.Links.GetByShortID(c.Params("id"))
	if err != nil {
		return c.JSONError(err)
	}
	if link.Password != "" {
		return c.JSONError(raptor.NewErrorUnauthorized("Link is password protected"))
	}
	return c.JSON(link.ToPublicLink())
}

func (lc *LinksController) Create(c *raptor.Context) error {
	var link models.Link
	if err := c.BodyParser(&link); err != nil {
		return c.JSONError(raptor.NewErrorBadRequest("Invalid JSON"))
	}
	l, err := lc.Links.Create(link)
	if err != nil {
		return c.JSONError(err)
	}
	return c.JSON(l.ToPublicLink(), http.StatusCreated)
}
