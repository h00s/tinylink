package controllers

import (
	"net/http"

	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/app/services"
)

type LinksController struct {
	raptor.Controller

	Links    *services.LinksService
	Accesses *services.AccessesService
}

func (lc *LinksController) Get(c *raptor.Context) error {
	link, err := lc.Links.GetByShortID(c.Params("id"))
	if err != nil {
		return c.JSONError(err)
	}
	if link.Password != "" {
		// TODO: redirect to app to enter password
		return c.SendStatus(http.StatusUnauthorized)
	}
	return c.JSON(link.ToPublicLink())
}

func (lc *LinksController) Redirect(c *raptor.Context) error {
	link, err := lc.Links.GetByShortID(c.Params("shortID"))
	if err != nil {
		return c.JSONError(err)
	}
	if link.Password != "" {
		return c.JSONError(raptor.NewErrorUnauthorized("Link is password protected"))
	}
	go lc.Accesses.Create(link.ID)
	return c.Redirect(link.URL)
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
