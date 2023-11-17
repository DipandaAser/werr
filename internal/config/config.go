package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	configFile = "./config/config.yaml"
)

type Config struct {
	MONGO_URI       string
	DATABASE_NAME   string
	GORSE_URL       string
	GORSE_APIKEY    string
	S3_ACCESS_KEY   string
	S3_SECRET_KEY   string
	S3_ENDPOINT     string
	S3_BUCKET       string
	S3_REGION       string
	ALLOWED_ORIGINS []string
}

var config *Config

func LoadConfiguration() error {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	config = &Config{
		MONGO_URI:       getConfigValueOrPanic("MONGO_URI"),
		DATABASE_NAME:   getConfigValueOrPanic("DATABASE_NAME"),
		GORSE_URL:       getConfigValueOrPanic("GORSE_URL"),
		GORSE_APIKEY:    viper.GetString("GORSE_APIKEY"),
		S3_ACCESS_KEY:   getConfigValueOrPanic("S3_ACCESS_KEY"),
		S3_SECRET_KEY:   getConfigValueOrPanic("S3_SECRET_KEY"),
		S3_ENDPOINT:     getConfigValueOrPanic("S3_ENDPOINT"),
		S3_BUCKET:       getConfigValueOrPanic("S3_BUCKET"),
		S3_REGION:       getConfigValueOrPanic("S3_REGION"),
		ALLOWED_ORIGINS: viper.GetStringSlice("ALLOWED_ORIGINS"),
	}

	return nil
}

func GetConfig() *Config {
	return config
}

func getConfigValueOrPanic(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Fatalf("%s MUST be provided in config.yaml file\n", key)
	}
	return value
}
