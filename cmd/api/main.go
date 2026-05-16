package main

import (
	"go-links/configs"
	"go-links/internal/database"
	"go-links/internal/server"
	"log"
)

func main() {
	secret := configs.GetSecret()
	db := database.NewDatabase(secret)

	srv := server.NewServer(db)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
