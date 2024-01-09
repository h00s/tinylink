package config

import "github.com/h00s/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
