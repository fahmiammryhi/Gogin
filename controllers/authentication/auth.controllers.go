package authentication

import (
	"Gogin/databases"
	"Gogin/models"
	"Gogin/request/authenticationRequest"
	"Gogin/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

// Login Handler
func Login(ctx *gin.Context) {
	var loginRequest authenticationRequest.LoginRequest

	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request",
			"error":   err.Error(),
		})
		return
	}

	validate := validator.New()
	if errValidate := validate.Struct(loginRequest); errValidate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
		return
	}

	var user models.User
	err := databases.DB.Table("adm_users").Where("email = ?", loginRequest.Email).First(&user).Error

	if err != nil {
		log.Printf("Login failed for Email: %s, user not found.", loginRequest.Email)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid Email",
		})
		return
	}

	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		log.Printf("Login failed for email: %s, invalid password.", loginRequest.Email)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid password",
		})
		return
	}

	// Buat JWT claims untuk access token
	claims := jwt.MapClaims{
		"id_user":  user.IDUser,
		"id_role":  user.IDRole,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Minute * 2).Unix(),
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Failed to generate token",
		})
		return
	}

	// Buat refresh token
	refreshClaims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 7).Unix(),
	}

	refreshToken, err := utils.GenerateToken(&refreshClaims)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error generating refresh token",
		})
		return
	}

	// Set cookies
	ctx.SetCookie(
		"access_token",
		token,
		int(time.Minute*2/time.Second),
		"/",
		"",
		false,
		true,
	)

	ctx.SetCookie(
		"refresh_token",
		refreshToken,
		int(time.Hour*7/time.Second),
		"/",
		"",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}

// Refresh Access Token Handler
func RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Refresh token not found",
		})
		return
	}

	newAccessToken, err := RefreshAccessToken(refreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Set new access token cookie
	ctx.SetCookie(
		"access_token",
		newAccessToken,
		int(time.Minute*2/time.Second),
		"/",
		"",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message":      "Token refreshed successfully",
		"access_token": newAccessToken,
	})
}

// Helper function untuk refresh access token
func RefreshAccessToken(refreshToken string) (string, error) {
	claims, err := utils.DecodeToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid or expired refresh token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", errors.New("invalid email in refresh token")
	}

	var user models.User
	err = databases.DB.Table("adm_users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", errors.New("user not found")
	}

	// Buat claims baru untuk access token
	newClaims := jwt.MapClaims{
		"id_user":  user.IDUser,
		"id_role":  user.IDRole,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Minute * 2).Unix(),
	}

	newAccessToken, err := utils.GenerateToken(&newClaims)
	if err != nil {
		return "", errors.New("failed to generate new access token")
	}

	return newAccessToken, nil
}

// Logout Handler
func Logout(ctx *gin.Context) {
	accessToken, err1 := ctx.Cookie("access_token")
	refreshToken, err2 := ctx.Cookie("refresh_token")

	if (err1 != nil || accessToken == "") && (err2 != nil || refreshToken == "") {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Anda belum login",
		})
		return
	}

	// Hapus access token
	ctx.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	// Hapus refresh token
	ctx.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout berhasil",
	})
}
