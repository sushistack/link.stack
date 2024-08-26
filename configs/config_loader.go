package configs

import (
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/sushistack/link.stack/utils"

	"github.com/spf13/viper"
)

// Config (= config.yml)
type Config struct {
	App struct {
		Name string `mapstructure:"name"`
	}
	Datasource *Datasource
}

type Datasource struct {
	URI            string `mapstructure:"uri"`
	Username       string `mapstructure:"username"`
	Password       string `mapstructure:"password"`
	DatabaseName   string `mapstructure:"db"`
	ConnectionPool struct {
		MinSize uint64 `mapstructure:"min"`
		MaxSize uint64 `mapstructure:"max"`
		MaxIdle int    `mapstructure:"max"`
	} `mapstructure:"connection-pool"`
}

func NewDefaultConfig() *Config {
	return &Config{}
}

type ConfigOptions struct {
	ConfigFilePath string
}

func DefaultOptions() *ConfigOptions {
	return &ConfigOptions{
		ConfigFilePath: utils.ProjectRoot + "/configs/config.yaml",
	}
}

func LoadConfig(options *ConfigOptions) *Config {
	if options == nil {
		options = DefaultOptions()
	}
	viper.SetConfigFile(options.ConfigFilePath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		utils.Logger.WithFields(logrus.Fields{
			"filePath": options.ConfigFilePath,
		}).Error("Can not read config file.", err)
		return NewDefaultConfig()
	}

	env := LoadEnvironment(nil)
	config := viper.AllSettings()
	replaceEnvVariables(config, env)

	var cfg Config
	if err := mapstructure.Decode(config, &cfg); err != nil {
		utils.Logger.WithFields(logrus.Fields{
			"config": config,
		}).Error("Can not decode config map.", err)
		return NewDefaultConfig()
	}

	return &cfg
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
