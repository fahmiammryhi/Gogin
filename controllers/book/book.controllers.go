package book

import(
	"github.com/gin-gonic/gin"
)
func GetAllBooks(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Get all books",
	})

}