package configs

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (config AppConfig, err error) {

	fmt.Println(os.Getenv("APP_ENV"))
	godotenv.Load()
	if os.Getenv("APP_ENV") == "local" {
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("HTTP_PORT not set in environment variables")
	}

	fmt.Println("Configuration loaded: ", httpPort)
	return AppConfig{ServerPort: httpPort}, nil
}
