package initializers

import (
	"github.com/go-raptor/connector/postgres"
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/db"
)

func Database() raptor.Database {
	return raptor.Database{
		Connector:  postgres.NewPostgresConnector(),
		Migrations: db.Migrations(),
	}
}
