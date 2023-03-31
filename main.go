package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/paulomujuru/simplebank/api"
	db "github.com/paulomujuru/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"                                                            // TODO: read from env
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" // TODO: read from env
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(":8080"); err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
