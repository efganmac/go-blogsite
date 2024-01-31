package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (conf AppConfig, err error) {
	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}

	LaunchPort := os.Getenv("LOCAL_HOST")
	if len(LaunchPort) < 1 {
		return AppConfig{}, errors.New("Your http port is wrong, please check!")
	}

	return AppConfig{ServerPort: LaunchPort}, nil
}
