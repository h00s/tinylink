package config

import "github.com/go-raptor/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Scope("/links",
				raptor.Route("POST", "", "LinksController", "Create"),
			),
		),
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
