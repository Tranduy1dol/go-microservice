package main

import (
	"github.com/Tranduy1dol/go-microservice/internal/db"
	"github.com/Tranduy1dol/go-microservice/internal/env"
	"github.com/Tranduy1dol/go-microservice/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		dbConfig: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:password@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env:     env.GetString("ENV", "development"),
		version: env.GetString("VERSION", "0.0.1"),
	}

	db, err := db.New(
		cfg.dbConfig.addr,
		cfg.dbConfig.maxOpenConns,
		cfg.dbConfig.maxIdleConns,
		cfg.dbConfig.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("Connected to database")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
