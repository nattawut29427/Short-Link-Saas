package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"

	"go-links/configs"
	"go-links/internal/database"
	"go-links/internal/server"
)

func main() {
	secret := configs.GetSecret()
	config := configs.LoadConfig()

	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	db := database.NewDatabase(secret)

	rdb := database.NewRedisClient(secret)

	srv := server.NewServer(config.App.Port, db, rdb, secret.JWTSecret)

	fmt.Printf("server listening on port %d\n", config.App.Port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
