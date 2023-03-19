package util

import (
	"github.com/spf13/viper"
)

const (
	Development = "development"
	Production  = "production"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment          string `mapstructure:"ENVIRONMENT"`
	DBDriver             string `mapstructure:"DB_DRIVER"`
	DBSource             string `mapstructure:"DB_SOURCE"`
	Token                string `mapstructure:"TOKEN"`
	SandboxToken         string `mapstructure:"SANDBOX_TOKEN"`
	URLhistoricalCandles string `mapstructure:"URL_HISTORICAL_CANDLES"`
	URLAPI               string `mapstructure:"URL_API"`
	SandboxURLAPI        string `mapstructure:"SANDBOX_URL_API"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
