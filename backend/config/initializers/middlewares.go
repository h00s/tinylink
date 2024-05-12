package initializers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/middlewares"
)

func Middlewares() raptor.Middlewares {
	return raptor.Middlewares{
		raptor.Use(middlewares.NewLimiterMiddleware()),
	}
}
