package mysql

import (
	"errors"
	"fmt"
	"log"
	"microblog/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(cfg *config.Config) {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New(
		"file://db/mysql/migrations",
		fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s", cfg.DbUsername, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbDatabase),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
}
