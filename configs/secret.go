package configs

import (
	"log"
	"github.com/spf13/viper"
)

func GetSecret() *Secrets {
	var secret Secrets

	v := viper.New()
	v.AddConfigPath("./configs")
	v.SetConfigName("secret")
	v.SetConfigType("json")

	// Enable environment variables fallback/override
	v.AutomaticEnv()

	// Explicitly bind struct fields to environment variables
	v.BindEnv("postgres_host", "POSTGRES_HOST")
	v.BindEnv("postgres_port", "POSTGRES_PORT")
	v.BindEnv("postgres_user", "POSTGRES_USER")
	v.BindEnv("postgres_password", "POSTGRES_PASSWORD")
	v.BindEnv("postgres_db", "POSTGRES_DB")

	v.BindEnv("redis_host", "REDIS_HOST")
	v.BindEnv("redis_port", "REDIS_PORT")
	v.BindEnv("redis_password", "REDIS_PASSWORD")
	v.BindEnv("redis_db", "REDIS_DB")

	v.BindEnv("jwt_secret", "JWT_SECRET")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Warning: Config file 'secret.json' not found (relying on Environment Variables): %s", err)
	}

	if err := v.Unmarshal(&secret); err != nil {
		log.Fatalf("Failed to unmarshal secret file: %s", err)
	}

	return &secret
}
