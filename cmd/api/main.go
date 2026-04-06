package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/akuumaa/nexo/internal/config"
	"github.com/akuumaa/nexo/internal/db"
	"github.com/akuumaa/nexo/internal/http/router"
)

func main() {
	cfg := config.Load()

	database, err := db.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	r := router.New(database)

	addr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	log.Printf("starting nexo in %s on %s", cfg.AppEnv, addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
