package utils

import (
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

func LoadConfig(configFilePath string) map[string]interface{} {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {

		return make(map[string]interface{})
	}

	env := LoadEnvironment(nil)
	config := viper.AllSettings()
	replaceEnvVariables(config, env)

	return config
}

func replaceEnvVariables(config map[string]interface{}, env map[string]string) {
	for key, value := range config {
		switch v := value.(type) {
		case string:
			config[key] = getEnvValue(env, v)
		case map[string]interface{}:
			replaceEnvVariables(v, env)
		}
	}
}

func getEnvValue(env map[string]string, value string) string {
	if strings.HasPrefix(value, "{{") && strings.HasSuffix(value, "}}") {
		return env[ExtractBetweenBraces(value)]
	}
	return value
}

func ExtractBetweenBraces(s string) string {
	start := strings.Index(s, "{{")
	if start == -1 {
		return s
	}

	start += len("{{")
	end := strings.Index(s[start:], "}}")
	if end == -1 {
		return s
	}

	return strings.TrimSpace(s[start : start+end])
}
