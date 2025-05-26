package main

import (
	"github.com/Tranduy1dol/go-microservice/internal/db"
	"github.com/Tranduy1dol/go-microservice/internal/env"
	store2 "github.com/Tranduy1dol/go-microservice/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://user:password@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 30, 30, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store2.NewStorage(conn)
	db.Seed(store)
}
