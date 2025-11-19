package utils

import (
	"Gogin/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(ctx *gin.Context) (page, pageSize, offset int) {
	page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if page < 1 {
		page = 1
	}

	pageSize, _ = strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	if pageSize < 1 {
		pageSize = 10
	}

	offset = (page - 1) * pageSize
	return
}

func BuildPaginationResponse(page, pageSize int, total int64) response.ResponsePage {
	return response.ResponsePage{
		Page:       uint(page),
		PageSize:   uint(pageSize),
		TotalItems: uint(total),
		TotalPages: uint((total + int64(pageSize) - 1) / int64(pageSize)),
	}
}

func SuccessResponse(ctx *gin.Context, data interface{}, pagination *response.ResponsePage) {
	resp := response.Response{
		StatusCode: "200",
		Message:    "success",
		Data:       data,
	}

	if pagination != nil {
		resp.Pagination = pagination
	}

	ctx.JSON(200, resp)
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"status_code": strconv.Itoa(statusCode),
		"message":     message,
	})
}
