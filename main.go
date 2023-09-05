package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mikcatta/simple_bank/api"

	db "github.com/mikcatta/simple_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:123456@192.168.1.83:5432/simplebank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the db:", err)
	}

	store := db.NewStore(conn)

	fmt.Println("Starting .... ", conn.Stats())

	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
