package bootstrap

import (
	"github.com/godazz/bloGin/internal/database/seeder"
	"github.com/godazz/bloGin/pkg/config"
	"github.com/godazz/bloGin/pkg/db"
)

func Seed() {
	config.Set()

	db.Connect()

	seeder.Seed()
}
