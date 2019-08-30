package model

import (
	"testing"
)

func TestComparePassword(t *testing.T) {
	var user = User{
		Password: "good_password",
	}

	user.HashPassword()

	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"password is valid", args{"good_password"}, true},
		{"password is invalid", args{"bad_password"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := user.ComparePassword(tt.args.password); got != tt.want {
				t.Errorf("ComparePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
