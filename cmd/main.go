package main

import "github.com/HarAlb/go-project/internal/api"

func main() {
	api := api.NewAPI()

	api.Start()
}