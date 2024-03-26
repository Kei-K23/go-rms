package main

import (
	"log"

	"github.com/Kei-K23/go-rms/backend/cmd/api"
	"github.com/Kei-K23/go-rms/backend/internal/config"
	"github.com/Kei-K23/go-rms/backend/internal/db"
)

func main() {

	ser := api.NewAPIServer(":4000", nil)

	sqlDB, err := db.NewDB(config.Env.DB_CONNECTION_STRING)

	if err != nil {
		log.Fatal(err)
	}

	db.InitDB(sqlDB)

	ser.Run()
}
