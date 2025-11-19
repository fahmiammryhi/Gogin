package utils

import (
	"errors"
	"log"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// HashingPassword meng-hash password menggunakan bcrypt
func HashingPassword(password string) (string, error) {
	// Validasi password sebelum di-hash
	if err := validatePassword(password); err != nil {
		return "", err // Kembalikan error jika password tidak valid
	}

	// Hash the password with bcrypt and a cost factor of 14
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err // Return an empty string and the error if hashing fails
	}
	return string(hashedByte), nil // Return the hashed password as a string
}

// CheckPasswordHash membandingkan password plaintext dengan hashed password
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Printf("Password mismatch: %v", err)
	}
	return err == nil
}

// validatePassword memvalidasi password berdasarkan aturan tertentu
// Password harus minimal 6 karakter, 1 huruf besar, 1 angka, dan 1 karakter spesial
func validatePassword(password string) error {
	// Password length check (minimum 6 characters)
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	// Check if there's at least one uppercase letter
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUppercase {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check if there's at least one number
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasNumber {
		return errors.New("password must contain at least one number")
	}

	// Check if there's at least one special character
	hasSpecialChar := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-\=\{\}\[\]:;"'<>,\./?\\|~]`).MatchString(password)

	if !hasSpecialChar {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

// Decode dan validasi refresh token
func ValidateRefreshToken(tokenString string) (*jwt.RegisteredClaims, error) {
	// Validasi token string tidak kosong
	if tokenString == "" {
		return nil, errors.New("refresh token is empty")
	}

	// Parse token dengan claims
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validasi signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SecretKey, nil
	})

	// Handle error dari parsing
	if err != nil {
		// Cek tipe error spesifik
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("refresh token has expired")
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("malformed refresh token")
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("refresh token not valid yet")
		}
		return nil, errors.New("invalid refresh token: " + err.Error())
	}

	// Validasi token dan claims
	if token == nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Extract claims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Validasi expiration time
	if claims.ExpiresAt == nil {
		return nil, errors.New("token missing expiration time")
	}

	// Cek apakah refresh token sudah kadaluarsa
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("refresh token has expired")
	}

	// Validasi issued at time (opsional tapi direkomendasikan)
	if claims.IssuedAt != nil && claims.IssuedAt.Time.After(time.Now()) {
		return nil, errors.New("token used before issued")
	}

	// Validasi not before time (opsional)
	if claims.NotBefore != nil && claims.NotBefore.Time.After(time.Now()) {
		return nil, errors.New("token not valid yet")
	}

	return claims, nil
}
