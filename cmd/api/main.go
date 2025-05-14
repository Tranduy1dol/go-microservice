package main

import (
	"github.com/Tranduy1dol/go-microservice/internal/env"
	"github.com/Tranduy1dol/go-microservice/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
