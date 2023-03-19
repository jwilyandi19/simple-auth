package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("[LoadConfig] Failed to read config")
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("[LoadConfig] Failed to unmarshal config")
		return Config{}, err
	}

	return config, err
}
