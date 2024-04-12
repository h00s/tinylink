package db

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/tinylink/db/migrate"
)

func Migrations() raptor.Migrations {
	return raptor.Migrations{
		1: migrate.AddPaste,
	}
}
