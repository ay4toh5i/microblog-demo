package main

import (
	"fmt"
	"log"
	"microblog/app/web"
	"microblog/config"
	"microblog/db/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	log.Println("server start")

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DbUsername,
			cfg.DbPassword,
			cfg.DbHost,
			cfg.DbPort,
			cfg.DbDatabase,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	mysql.Migrate(cfg)

	web.RunServer(db)(cfg.HttpPort)
}
