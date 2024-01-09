package main

import (
	"github.com/h00s/raptor"
	"github.com/h00s/tinylink/config"
	"github.com/h00s/tinylink/config/initializers"
)

func main() {
	r := raptor.NewRaptor()

	r.Init(initializers.App())
	r.Routes(config.Routes())

	r.Listen()
}
