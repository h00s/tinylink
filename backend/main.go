package main

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/config"
	"github.com/h00s/tinylink/config/initializers"
)

func main() {
	r := raptor.NewRaptor(
		initializers.App(),
		config.Routes(),
	)

	r.Listen()
}
