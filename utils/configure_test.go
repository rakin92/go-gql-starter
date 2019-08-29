package utils

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigureEnv(t *testing.T) {
	ConfigureEnv()
	t.Parallel()

	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	env := viper.GetString("ENV")

	assert.Equal(t, env, "LOCAL", "environment is local by default")
	assert.Equal(t, dbHost, "localhost", "db host should be local by default")
	assert.Equal(t, dbPort, "5432", "db port should be 5432 by default")
	assert.Equal(t, dbName, "local-db", "db name should be local-db by default")
}
