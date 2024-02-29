package initializers

import (
	"github.com/go-raptor/raptor"
)

func App() *raptor.AppInitializer {
	return &raptor.AppInitializer{
		Database:    Database(),
		Services:    Services(),
		Middlewares: Middlewares(),
		Controllers: Controllers(),
	}
}
