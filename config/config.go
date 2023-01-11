package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
}

type ConfigEnvContract interface {
	Get(string) any
	GetString(string) string
}

func LoadEnv() ConfigEnvContract {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &ConfigEnv{}
}

func (c *ConfigEnv) Get(env string) any {
	data := os.Getenv(env)

	return data
}

func (c *ConfigEnv) GetString(env string) string {
	data := c.Get(env).(string)
	return data
}
