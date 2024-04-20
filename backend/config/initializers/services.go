package initializers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/services"
)

func Services() raptor.Services {
	return raptor.Services{
		&services.LinksService{},
		&services.AccessesService{},
	}
}
