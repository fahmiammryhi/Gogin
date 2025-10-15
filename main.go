package main

import "github.com/gin-gonic/gin"

func main() {

	app := gin.Default()
	route := app
	route.GET("/", func(ctx *gin.Context) {
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
	})

	app.Run(":8080")
}
