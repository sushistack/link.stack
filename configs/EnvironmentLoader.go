package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sushistack/link.stack/utils"
)

type EnvironmentOptions struct {
	EnvFilePath string
}

func DefaultConfigOptions() *EnvironmentOptions {
	return &EnvironmentOptions{
		EnvFilePath: utils.ProjectRoot + "/configs/" + ".env",
	}
}

func LoadEnvironment(options *EnvironmentOptions) map[string]string {
	if options == nil {
		options = DefaultConfigOptions()
	}

	if err := godotenv.Load(options.EnvFilePath); err != nil {
		return make(map[string]string)
	}

	environment := map[string]string{
		"MONGODB_URI":      getEnv("MONGODB_URI", ""),
		"MONGODB_USERNAME": getEnv("MONGODB_USERNAME", ""),
		"MONGODB_PASSWORD": getEnv("MONGODB_PASSWORD", ""),
	}

	return environment
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
