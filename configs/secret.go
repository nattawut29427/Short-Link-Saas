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

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read secret file: %s", err)
	}

	if err := v.Unmarshal(&secret); err != nil {
		log.Fatalf("Failed to unmarshal secret file: %s", err)
	}

	return &secret
}
