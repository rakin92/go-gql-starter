package utils

import (
	"testing"
)

func TestValidateJWT(t *testing.T) {
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDE5LTA0LTIyVDIyOjQ3OjAxLjA1NTU5MSswOTowMCIsInVzZXJJRCI6IjEifQ.3dOqg_ceqv25nCl2C2WL_lye6vLC5dE8nplk184-5lQ"

	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"fails on invalid token", args{invalidToken}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := ValidateJWT(&tt.args.token); got != tt.want {
				t.Errorf("ValidateJWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
