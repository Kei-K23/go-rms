package main

import "github.com/Kei-K23/go-rms/backend/cmd/api"

func main() {

	ser := api.NewAPIServer(":4000", nil)

	ser.Run()
}
