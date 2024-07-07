package utils

import (
	"testing"
)
var password string ="password 123"

func TestHashPassword(t *testing.T) {
	hashedPassword,err:= HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if hashedPassword == "" {
		t.Fatalf("Expected a hashed password, got an empty string")
	}
}
func TestCheckPasswordHash(t *testing.T) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !CheckPasswordHash(password, hashedPassword) {
		t.Fatalf("Expected passwords to match")
	}

	wrongPassword := "wrongpassword"
	if CheckPasswordHash(wrongPassword, hashedPassword) {
		t.Fatalf("Expected passwords to not match")
	}
}