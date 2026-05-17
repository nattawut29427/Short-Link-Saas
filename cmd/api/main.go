package main

import (
	"fmt"
	"log"

	"go-links/configs"
	"go-links/internal/database"
	"go-links/internal/server"
)

func main() {
	secret := configs.GetSecret()
	config := configs.LoadConfig()
	db := database.NewDatabase(secret)

	rdb := database.NewRedisClient(secret)

	srv := server.NewServer(config.App.Port, db, rdb)

	fmt.Printf("server listening on port %d\n", config.App.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
