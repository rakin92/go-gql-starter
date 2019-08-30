package migrations

import (
	"github.com/rakin92/go-gql-starter/db"
	"github.com/rakin92/go-gql-starter/model"
)

// MigrateDB : auto migrates the given database
func MigrateDB(db *db.DB) {
	db.AutoMigrate(&model.User{})
}
