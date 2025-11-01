package main

import (
	"backend-app/configs"
	"backend-app/internal/api"
	"fmt"
)

func main() {
	fmt.Println("Starting the backend application...")
	config, err := configs.SetupEnv()
	if err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}
	api.StartServer(config)

}
