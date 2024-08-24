package utils

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config (= config.yml)
type Config struct {
	App struct {
		Name string `mapstructure:"name"`
	}
	Datasource struct {
		URI      string `mapstructure:"uri"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
}

func LoadConfig(configFilePath string) (map[string]interface{}, error) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	config := viper.AllSettings()
	replaceEnvVariables(config)

	return config, nil
}

func replaceEnvVariables(config map[string]interface{}) {
	env, _ := LoadEnvironment(nil)

	for key, value := range config {
		switch v := value.(type) {
		case string:
			config[key] = getEnvValue(env, v)
		case map[string]interface{}:
			replaceEnvVariables(v)
		}
	}
}

func getEnvValue(env map[string]string, value string) string {
	if strings.HasPrefix(value, "{{") && strings.HasSuffix(value, "}}") {

		value, ok := env[ExtractBetweenBraces(value)]
		if !ok {
			return value
		}
	}
	return value
}

func ExtractBetweenBraces(s string) string {
	start := strings.Index(s, "{{")
	if start == -1 {
		return ""
	}

	start += len("{{")
	end := strings.Index(s[start:], "}}")
	if end == -1 {
		return ""
	}

	return s[start : start+end]
}
