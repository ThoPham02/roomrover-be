package utils_test

import (
	"testing"

	"roomrover/service/account/utils"
)

func TestHashPassword_NonEmptyAndDifferentLength(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{name: "Test Case 1", password: "password123"},
		{name: "Test Case 2", password: "abc123"},
		{name: "Test Case 3", password: "1234567890"},
		{name: "Test Case 4", password: "special@#"},
		{name: "Test Case 5", password: "longpasswordwithmorethan20characters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := utils.HashPassword(tt.password)
			if err != nil {
				t.Fatalf("Error hashing password: %v", err)
			}

			if len(hashedPassword) == 0 {
				t.Error("Expected non-empty hashed password")
			}

			if len(hashedPassword) == len(tt.password) {
				t.Error("Expected hashed password to have a different length than the original password")
			}
		})
	}
}
