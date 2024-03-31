package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_CONNECTION_STRING string
	SECRET_KEY           string
	STRIPE_API_KEY       string
}

var Env Config = initEnv()

func initEnv() Config {
	godotenv.Load()

	return Config{
		DB_CONNECTION_STRING: getEnv("DB_CONNECTION_STRING", "dbeaver:dbeaver@(127.0.0.1:3306)/go-rms?multiStatements=true"),
		SECRET_KEY:           getEnv("SECRET_KEY", "this-is-secret-key"),
		STRIPE_API_KEY:       getEnv("STRIPE_API_KEY", "this-is-secret-key"),
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
