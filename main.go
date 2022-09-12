package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/thinhngtruong/go-banking/api"
	db "github.com/thinhngtruong/go-banking/db/sqlc"
	"github.com/thinhngtruong/go-banking/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatalln("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatalln(err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
