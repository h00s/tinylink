package migrate

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
)

func AddAccess(db *raptor.DB) error {
	return db.Migrator().CreateTable(&models.Access{})
}
