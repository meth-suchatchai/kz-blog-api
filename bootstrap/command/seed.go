package command

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/kuroshibaz/config"
	"github.com/kuroshibaz/lib/gormdb"
)

func Seed(cfg *config.Env) {
	db, err := gormdb.ConnectSQL(&cfg.Database)
	if err != nil {
		log.Fatalf("error connect SQL: %v", err)
	}

	err = db.Migrate()
	//db.Seed()
	if err != nil {
		log.Fatalf("error migrate: %v", err)
	}

	db.Seed()
}
