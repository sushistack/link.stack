package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentOptions struct {
	EnvFilePath string
}

func DefaultConfigOptions() *EnvironmentOptions {
	return &EnvironmentOptions{
		EnvFilePath: "../configs/.env",
	}
}

func LoadEnvironment(options *EnvironmentOptions) (map[string]string, error) {
	if options == nil {
		options = DefaultConfigOptions()
	}

	if err := godotenv.Load(options.EnvFilePath); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}

	environment := map[string]string{
		"MONGODB_URI":      getEnv("MONGODB_URI", ""),
		"MONGODB_USERNAME": getEnv("MONGODB_USERNAME", ""),
		"MONGODB_PASSWORD": getEnv("MONGODB_PASSWORD", ""),
	}

	return environment, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
