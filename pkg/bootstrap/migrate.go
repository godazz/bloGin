package bootstrap

import (
	"github.com/godazz/bloGin/internal/database/migration"
	"github.com/godazz/bloGin/pkg/config"
	"github.com/godazz/bloGin/pkg/db"
)

func Migrate() {
	config.Set()

	db.Connect()

	migration.Migrate()
}
