package main

import (
	"go-api/configs"
	"go-api/internal/database"
	"go-api/internal/server"
	"log"
)

// @title Go API
// @version 1.0
// @description Boilerplate Go API
// @host localhost:8080
// @basePath /v1
func main() {
	secret := configs.GetSecret()
	db := database.NewDatabase(secret)

	srv := server.NewServer(db)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
