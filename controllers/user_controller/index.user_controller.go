package user_controller

import (
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

	ctx.JSON(200, gin.H{
		"message": "Get all users",
	})
}
