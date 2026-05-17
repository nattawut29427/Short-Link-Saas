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
	v.BindEnv("mysql_host", "MYSQL_HOST")
	v.BindEnv("mysql_port", "MYSQL_PORT")
	v.BindEnv("mysql_user", "MYSQL_USER")
	v.BindEnv("mysql_password", "MYSQL_PASSWORD")
	v.BindEnv("mysql_db", "MYSQL_DB")

	v.BindEnv("redis_host", "REDIS_HOST")
	v.BindEnv("redis_port", "REDIS_PORT")
	v.BindEnv("redis_password", "REDIS_PASSWORD")
	v.BindEnv("redis_db", "REDIS_DB")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Warning: Config file 'secret.json' not found (relying on Environment Variables): %s", err)
	}

	if err := v.Unmarshal(&secret); err != nil {
		log.Fatalf("Failed to unmarshal secret file: %s", err)
	}

	return &secret
}
