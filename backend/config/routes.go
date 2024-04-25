package config

import "github.com/go-raptor/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1",
			raptor.Scope("/links",
				raptor.Route("POST", "/:id/authorize", "LinksController", "Authorize"),
				raptor.Route("GET", "/:id", "LinksController", "Get"),
				raptor.Route("POST", "", "LinksController", "Create"),
			),
		),
		raptor.Route("GET", "/links", "SPAController", "Index"),
		raptor.Route("GET", "/:shortID", "LinksController", "Redirect"),
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
