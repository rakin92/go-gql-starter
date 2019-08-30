package utils

import (
	"testing"

	"github.com/spf13/viper"
)

func TestConfigureEnv(t *testing.T) {
	ConfigureEnv()

	type args struct {
		envKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"fails on invalid token", args{"DB_HOST"}, "localhost"},
		{"fails on invalid token", args{"DB_PORT"}, "5432"},
		{"fails on invalid token", args{"DB_NAME"}, "local-db"},
		{"fails on invalid token", args{"ENV"}, "LOCAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := viper.GetString(tt.args.envKey); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
