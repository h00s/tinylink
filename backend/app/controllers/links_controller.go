package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
	"github.com/h00s/tinylink/app/services"
	"github.com/h00s/tinylink/internal"
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
		return c.Redirect(fmt.Sprintf("/links/%s/authorize", c.Params("shortID")))
	}
	go lc.Accesses.Create(link.ID)
	return c.Redirect(link.URL)
}

func (lc *LinksController) Authorize(c *raptor.Context) error {
	l, err := lc.Links.GetByShortID(c.Params("id"))
	if err != nil {
		return c.JSONError(err)
	}
	var link models.PublicLink
	if err := c.BodyParser(&link); err != nil {
		return c.JSONError(raptor.NewErrorBadRequest("Invalid JSON"))
	}
	if internal.ComparePasswords(l.Password, link.Password) != nil {
		return c.JSONError(raptor.NewErrorUnauthorized("Invalid password"))
	}
	return c.JSON(l.ToPublicLink())
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
