package routes

import (
	bookcontroller "Gogin/controllers/book_controller"
	"Gogin/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app
	route.GET("/", user_controller.GetTest)
	route.GET("/user", user_controller.GetAllUsers)
	route.GET("/book", bookcontroller.GetAllBooks)
}
