package configs

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port int `mapstructure:"port"`
		Name string `mapstructure:"name"`
	} `mapstructure:"app"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return &config
}
