package configs

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
}

func SetupEnv() (config AppConfig, err error) {

	fmt.Println(os.Getenv("APP_ENV"))
	godotenv.Load()
	if os.Getenv("APP_ENV") == "local" {
	}

	httpPort := os.Getenv("HTTP_PORT")
	dsn := os.Getenv("DATABASE_URL")

	if httpPort == "" {
		return AppConfig{}, errors.New("HTTP_PORT not set in environment variables")
	}

	if dsn == "" {
		return AppConfig{}, errors.New("DATABASE_URL not set in environment variables")
	}

	fmt.Println("Configuration loaded: ", httpPort)
	return AppConfig{ServerPort: httpPort, Dsn: dsn}, nil
}
