// internal/platform/config/config.go
package config

import (
	"github.com/spf13/viper"
	"imagine-proto/internal/llm/provider"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	LLM    LLMConfig    `mapstructure:"llm"`
	Redis  RedisConfig  `mapstructure:"redis"`
}

type ServerConfig struct {
	Port    string `mapstructure:"port"`
	Timeout int    `mapstructure:"timeout"`
}

type LLMConfig struct {
	Type        provider.ProviderType `mapstructure:"type"`
	BaseURL     string                `mapstructure:"baseUrl"`
	Model       string                `mapstructure:"model"`
	MaxTokens   int                   `mapstructure:"maxTokens"`
	Temperature float64               `mapstructure:"temperature"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func Load() (*Config, error) {
	viper.SetConfigName("dev")  // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
