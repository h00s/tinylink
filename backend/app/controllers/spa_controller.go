package controllers

import "github.com/h00s/raptor"

type SPAController struct {
	raptor.Controller
}

func (hc *SPAController) Index(c *raptor.Context) error {
	return c.SendFile("public/index.html")
}
