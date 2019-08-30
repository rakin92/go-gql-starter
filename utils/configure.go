package utils

import (
	"github.com/spf13/viper"
)

// ConfigureEnv sets up project environment
func ConfigureEnv() {
	viper.AutomaticEnv()

	viper.SetDefault("ENV", "LOCAL")

	// database
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "docker")
	viper.SetDefault("DB_PASSWORD", "docker")
	viper.SetDefault("DB_NAME", "local-db")
	viper.SetDefault("DB_SSL_MODE", "disable")
}
