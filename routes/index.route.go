package routes

import (
	"Gogin/controllers/administrator"
	"Gogin/controllers/authentication"
	"Gogin/controllers/book"
	"Gogin/controllers/test"
	"Gogin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoute(route *gin.Engine) {

	route.POST("/Login", authentication.Login)
	route.POST("/User/create", administrator.UserControllersCreate)

	api := route.Group("", middleware.AuthMiddleware())
	{
		// authentication route
		api.POST("/Logout", authentication.Logout)

		// administrator route
		api.GET("/User", administrator.UserControllersRead)
		api.GET("/User/:id", administrator.UserControllersReadByID)
		api.POST("/User/update/:id", administrator.UserControllersUpdate)
		api.POST("/User/delete/:id", administrator.UserControllersDelete)

		// // administrator Roles
		// api.GET("/Roles", administrator.RolesControllersRead)
		// api.GET("/Roles/:id", administrator.RolesControllersReadByID)
		// api.POST("/Roles/create", administrator.RolesControllersCreate)
		// api.POST("/Roles/update/:id", administrator.RolesControllersUpdate)
		// api.POST("/Roles/delete/:id", administrator.RolesControllersDelete)

		// // administrator Menus
		// api.GET("/Menus", administrator.MenusControllersRead)
		// api.GET("/Menus/:id", administrator.MenusControllersReadByID)
		// api.POST("/Menus/create", administrator.MenusControllersCreate)
		// api.POST("/Menus/update/:id", administrator.MenusControllersUpdate)
		// api.POST("/Menus/delete/:id", administrator.MenusControllersDelete)

		// // administrator Menu Roles
		// api.GET("/MenuRoles", administrator.MenuRolesControllersRead)
		// api.GET("/MenuRoles/:id", administrator.MenuRolesControllersReadByID)
		// api.POST("/MenuRoles/create", administrator.MenuRolesControllersCreate)
		// api.POST("/MenuRoles/update/:id", administrator.MenuRolesControllersUpdate)
		// api.POST("/MenuRoles/delete/:id", administrator.MenuRolesControllersDelete)

		api.GET("/test", test.GetTest)
		api.GET("/book", book.GetAllBooks)
	}
}
