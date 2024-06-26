package utils
import (
	"CloudSystem/utils"
	"testing"
	"time"
	"github.com/golang-jwt/jwt/v5"

)
const secretKey = "supersecret"
func TestGenerateToken(t *testing.T) {
	email := "test@example.com"
	userId := int64(12345)

	tokenString, err := utils.GenerateToken(email, userId)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["email"] != email {
			t.Errorf("Expected email %v, got %v", email, claims["email"])
		}
		if claims["userId"] != float64(userId) { // jwt.MapClaims stores numbers as float64
			t.Errorf("Expected userId %v, got %v", userId, claims["userId"])
		}
		exp := int64(claims["exp"].(float64)) // Convert the expiration time to int64
		if exp <= time.Now().Unix() {
			t.Errorf("Expected token to be valid, but it is expired")
		}
	} else {
		t.Errorf("Token is invalid")
	}
}