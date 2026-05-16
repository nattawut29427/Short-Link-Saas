package main

import (
	"fmt"
	"go-links/configs"
	"go-links/internal/database"
	"go-links/internal/server"
	"log"
)

func main() {
	secret := configs.GetSecret()
	config := configs.LoadConfig()

	db := database.NewDatabase(secret)

	srv := server.NewServer(config.App.Port, db)

	fmt.Printf("server listening on port %d\n", config.App.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
