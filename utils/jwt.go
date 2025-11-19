package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var SecretKey string

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	SecretKey = os.Getenv("JWT_SECRET")

	if SecretKey == "" {
		log.Fatal("JWT_SECRET is empty. Check your .env file")
	} else {
		fmt.Println("JWT_SECRET loaded successfully:", SecretKey) // Debugging
	}
}

// GenerateToken membuat JWT token baru dengan claims yang diberikan
func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webtoken, nil
}

// VerifToken memverifikasi JWT token
func VerifToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		fmt.Println("Token verification failed:", err) // Debugging
		return nil, err
	}

	return token, nil
}

// DecodeToken mendecode JWT token dan mengembalikan claims
func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GetTokenExpiry mendapatkan waktu kadaluarsa token
func GetTokenExpiry(token string) time.Time {
	parsedToken, _, err := new(jwt.Parser).ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return time.Now()
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return time.Now()
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return time.Now()
	}

	return time.Unix(int64(exp), 0)
}
