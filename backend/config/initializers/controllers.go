package initializers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.LinksController{},
		&controllers.SPAController{},
	}
}
