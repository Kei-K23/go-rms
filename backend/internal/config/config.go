package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_CONNECTION_STRING string
}

var Env Config = initEnv()

func initEnv() Config {
	godotenv.Load()

	return Config{
		DB_CONNECTION_STRING: getEnv("DB_CONNECTION_STRING", "dbeaver:dbeaver@127.0.0.1:3306)/go-rms?multiStatements=true"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
