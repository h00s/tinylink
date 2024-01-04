package main

import (
	"github.com/h00s/raptor"
	"github.com/h00s/tinylink/app/config"
)

func main() {
	r := raptor.NewRaptor()

	r.Middlewares(config.Middlewares())
	r.Services(config.Services())
	r.Controllers(config.Controllers())
	r.Routes(config.Routes())

	r.Listen()
}
