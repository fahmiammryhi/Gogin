package middleware

import (
	"Gogin/controllers/authentication"
	"Gogin/databases"
	"Gogin/models"
	"Gogin/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware adalah middleware untuk autentikasi menggunakan JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Mengambil access token dari cookies
		token, err := ctx.Cookie("access_token")

		// Jika access token tidak ditemukan, coba gunakan refresh token
		if err != nil || token == "" {
			fmt.Println("Access token not found, trying refresh token...") // Debugging

			// Coba ambil refresh token dari cookies
			refreshToken, err := ctx.Cookie("refresh_token")
			if err != nil || refreshToken == "" {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "unauthenticated",
				})
				ctx.Abort()
				return
			}

			// Gunakan refresh token untuk mendapatkan access token baru
			newAccessToken, err := authentication.RefreshAccessToken(refreshToken)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Failed to refresh access token",
					"error":   err.Error(),
				})
				ctx.Abort()
				return
			}

			// Simpan access token baru ke cookies
			ctx.SetCookie(
				"access_token",
				newAccessToken,
				int(time.Until(utils.GetTokenExpiry(newAccessToken)).Seconds()),
				"/",
				"",
				false, // Secure - set true jika menggunakan HTTPS
				true,  // HttpOnly
			)

			// Gunakan token baru untuk pengecekan selanjutnya
			token = newAccessToken
		}

		// Memeriksa apakah token valid
		claims, err := utils.DecodeToken(token)
		if err != nil {
			// Jika token kadaluarsa, coba refresh token
			if err.Error() == "Token is expired" {
				refreshToken, err := ctx.Cookie("refresh_token")
				if err != nil || refreshToken == "" {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"message": "Refresh token is required",
					})
					ctx.Abort()
					return
				}

				// Validasi refresh token
				_, err = utils.ValidateRefreshToken(refreshToken)
				if err != nil {
					// Jika refresh token expired atau tidak valid, hapus cookies dan minta login ulang
					ctx.SetCookie(
						"refresh_token",
						"",
						-1,
						"/",
						"",
						false,
						true,
					)

					ctx.JSON(http.StatusUnauthorized, gin.H{
						"message": "Refresh token expired or invalid, please log in again",
					})
					ctx.Abort()
					return
				}

				// Jika refresh token valid, buat access token baru
				newAccessToken, err := authentication.RefreshAccessToken(refreshToken)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"message": "Failed to refresh access token",
						"error":   err.Error(),
					})
					ctx.Abort()
					return
				}

				// Set access token baru di cookies
				ctx.SetCookie(
					"access_token",
					newAccessToken,
					int(time.Until(utils.GetTokenExpiry(newAccessToken)).Seconds()),
					"/",
					"",
					false,
					true,
				)

				// Decode ulang token
				claims, err = utils.DecodeToken(newAccessToken)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, gin.H{
						"message": "Failed to decode new access token",
					})
					ctx.Abort()
					return
				}

			} else {
				// Jika error lain, anggap token tidak valid
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "unauthenticated",
				})
				ctx.Abort()
				return
			}
		}

		// Simpan informasi user di context
		ctx.Set("userInfo", claims)

		// Lanjutkan ke handler berikutnya
		ctx.Next()
	}
}

// CheckMenuPermission memeriksa permission user terhadap menu tertentu
func CheckMenuPermission(ctx *gin.Context, menu string, permissionType string) (bool, error) {
	var menuRole models.MenuRole
	var IDRole int

	userInfo, exists := ctx.Get("userInfo")
	if !exists {
		return false, fmt.Errorf("failed to retrieve user information")
	}

	claims, ok := userInfo.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("invalid user information format")
	}

	idRoleFloat, ok := claims["id_role"].(float64)
	if !ok {
		return false, fmt.Errorf("invalid role ID format")
	}

	IDRole = int(idRoleFloat)

	// Perbaikan: gunakan nama tabel yang benar
	err := databases.DB.Table("adm_menu_roles").First(&menuRole, "id_role = ? AND id_menu = ?", IDRole, menu).Error
	if err != nil {
		// Tambahkan log untuk debugging
		log.Printf("Permission check failed - Role: %d, Menu: %s, Error: %v", IDRole, menu, err)
		return false, fmt.Errorf("user does not have access or menu not found")
	}

	switch permissionType {
	case "isInsert":
		return menuRole.IsInsert, nil
	case "isEdit":
		return menuRole.IsEdit, nil
	case "isView":
		return menuRole.IsView, nil
	case "isDelete":
		return menuRole.IsDelete, nil
	case "isPrint":
		return menuRole.IsPrint, nil
	case "isApprove":
		return menuRole.IsApprove, nil
	default:
		return false, fmt.Errorf("invalid permission type")
	}
}
