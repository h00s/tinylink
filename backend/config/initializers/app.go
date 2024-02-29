package initializers

import (
	"github.com/go-raptor/raptor"
)

func App() *raptor.AppInitializer {
	return &raptor.AppInitializer{
		Services:    Services(),
		Middlewares: Middlewares(),
		Controllers: Controllers(),
	}
}
