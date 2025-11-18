package user_controller

import (
	"Gogin/database"
	"Gogin/models"
	"Gogin/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTest(ctx *gin.Context) {

	isValidated := true

	if !isValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message":   "bad request",
			"validated": isValidated,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message":   "Hello, World!",
		"validated": isValidated,
	})
}

func GetAllUsers(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	users := new([]models.User)
	query := database.DB.Table("users")

	var total int64
	query.Count(&total)

	err = query.Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"status_code": "500",
			"message":     "internal server error",
		})
		return
	}

	responsePage := response.ResponsePage{
		Page:       uint(page),
		PageSize:   uint(pageSize),
		TotalItems: uint(total),
		TotalPages: uint((total + int64(pageSize) - 1) / int64(pageSize)),
	}

	response := response.Response{
		StatusCode: "200",
		Message:    "success",
		Pagination: responsePage,
		Data:       users,
	}

	ctx.JSON(200, response)
}
