package initializers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/controllers"
)

func App() *raptor.AppInitializer {
	return &raptor.AppInitializer{
		Services: raptor.Services{},

		Middlewares: raptor.Middlewares{},

		Controllers: raptor.Controllers{
			&controllers.SPAController{},
		},
	}
}
