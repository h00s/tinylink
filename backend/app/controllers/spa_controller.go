package controllers

import "github.com/go-raptor/raptor"

type SPAController struct {
	raptor.Controller
}

func (hc *SPAController) Index(c *raptor.Context) error {
	return c.SendFile("public/index.html")
}
