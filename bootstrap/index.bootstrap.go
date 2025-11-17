package bootstrap

import (
	"Gogin/configs/app_config"
	"Gogin/database"
	"Gogin/routes"

	"github.com/gin-gonic/gin"
)

func BootstrapApp() {

	database.ConnectDatabase()

	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
