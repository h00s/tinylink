package migrate

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/app/models"
)

func AddPaste(db *raptor.DB) error {
	return db.Migrator().CreateTable(&models.Link{})
}
