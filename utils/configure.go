package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// ConfigureEnv sets up project environment
func ConfigureEnv() {
	viper.AutomaticEnv()

	log.SetFormatter(&log.JSONFormatter{})

	viper.SetDefault("ENV", "LOCAL")

	// database
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "docker")
	viper.SetDefault("DB_PASSWORD", "docker")
	viper.SetDefault("DB_NAME", "local-db")
	viper.SetDefault("DB_SSL_MODE", "disable")

	if viper.GetString("LOGTYPE") != "CONSOLE" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	if viper.IsSet("LOGLEVEL") {
		level := map[string]log.Level{
			"INFO":  log.InfoLevel,
			"DEBUG": log.DebugLevel,
			"WARN":  log.WarnLevel,
			"ERROR": log.ErrorLevel,
		}
		log.SetLevel(level[viper.GetString("LOGLEVEL")])
	}
}
