package administrator

import (
	"Gogin/databases"
	"Gogin/middleware"
	"Gogin/models"
	"Gogin/request/administratorRequest"
	"Gogin/response/administratorResponse"
	"Gogin/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserControllersRead(ctx *gin.Context) {
	canView, errView := middleware.CheckMenuPermission(ctx, "1.2.1.3", "isView")
	if errView != nil {
		log.Println("Error:", errView.Error())
		utils.ErrorResponse(ctx, http.StatusUnauthorized, errView.Error())
		return
	}
	if !canView {
		utils.ErrorResponse(ctx, http.StatusForbidden, "User tidak memiliki akses untuk melihat data User")
		return
	}
	page, pageSize, offset := utils.GetPaginationParams(ctx)

	var users []administratorResponse.UserResponse
	query := databases.DB.Table("adm_users")

	var total int64
	query.Count(&total)

	if err := query.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		utils.ErrorResponse(ctx, 500, "internal server error")
		return
	}

	pagination := utils.BuildPaginationResponse(page, pageSize, total)
	utils.SuccessResponse(ctx, users, &pagination)
}

func UserControllersReadByID(ctx *gin.Context) {
	canView, errView := middleware.CheckMenuPermission(ctx, "1.2.1.3", "isView")
	if errView != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, errView.Error())
		return
	}
	if !canView {
		utils.ErrorResponse(ctx, http.StatusForbidden, "User tidak memiliki akses untuk melihat data User")
		return
	}

	page, pageSize, offset := utils.GetPaginationParams(ctx)
	IDUser := ctx.Param("id")

	var users []administratorResponse.UserResponse
	query := databases.DB.Table("adm_users").Where("id_user = ?", IDUser)

	var total int64
	query.Count(&total)

	if err := query.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		utils.ErrorResponse(ctx, 500, "internal server error")
		return
	}

	pagination := utils.BuildPaginationResponse(page, pageSize, total)
	utils.SuccessResponse(ctx, users, &pagination)
}

func UserControllersCreate(ctx *gin.Context) {
	userRequest := new(administratorRequest.UserRequest)
	if err := ctx.ShouldBindJSON(userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	// Validasi request
	validate := validator.New()
	errValidate := validate.Struct(userRequest)
	if errValidate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"error":   errValidate.Error(),
		})
		return
	}

	var role models.Role
	err := databases.DB.Table("adm_roles").First(&role, "id_role = ?", userRequest.IDRole).Error
	if err != nil {
		log.Println("Error fetching role:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid role ID",
			"error":   err.Error(),
		})
		return
	}

	// Get last ID user
	var lastNumber int
	err = databases.DB.Raw(`
		SELECT COALESCE(MAX(id_user), 0)
		FROM adm_users
	`).Scan(&lastNumber).Error
	if err != nil {
		log.Println("Error fetching last user ID:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Database error",
			"error":   err.Error(),
		})
		return
	}

	log.Println("Last Number:", lastNumber)
	newIdUser := lastNumber + 1

	// Hash password
	hashedPassword, err := utils.HashingPassword(userRequest.Password)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
		return
	}

	newUser := models.User{
		IDUser:              uint(newIdUser),
		IDRole:              userRequest.IDRole,
		Username:            userRequest.Username,
		Email:               userRequest.Email,
		Password:            hashedPassword,
		RefreshToken:        userRequest.RefreshToken,
		RefreshTokenExpired: userRequest.RefreshTokenExpired,
		IsActive:            true,
		CreatedDate:         time.Now(),
		UpdatedDate:         time.Now(),
	}

	errCreateUser := databases.DB.Table("adm_users").Create(&newUser).Error
	if errCreateUser != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to store data",
			"error":   errCreateUser.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    newUser,
	})
}

func UserControllersUpdate(ctx *gin.Context) {
	canEdit, errEdit := middleware.CheckMenuPermission(ctx, "1.2.1.1", "isEdit")
	if errEdit != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, errEdit.Error())
		return
	}
	if !canEdit {
		utils.ErrorResponse(ctx, http.StatusForbidden, "User tidak memiliki akses untuk mengubah data User")
		return
	}

	// ... kode update Anda
}

func UserControllersDelete(ctx *gin.Context) {
	canDelete, errDelete := middleware.CheckMenuPermission(ctx, "1.2.1.1", "isDelete")
	if errDelete != nil {
		utils.ErrorResponse(ctx, http.StatusUnauthorized, errDelete.Error())
		return
	}
	if !canDelete {
		utils.ErrorResponse(ctx, http.StatusForbidden, "User tidak memiliki akses untuk menghapus data User")
		return
	}

	// ... kode delete Anda
}
