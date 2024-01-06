package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/tinylink/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.SPAController{},
	}
}
