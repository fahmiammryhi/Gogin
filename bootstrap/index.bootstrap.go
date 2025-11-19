package bootstrap

import (
	"Gogin/configs"
	"Gogin/configs/app_config"
	"Gogin/databases"
	"Gogin/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BootstrapApp() {

	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// init config
	configs.InitConfig()

	// database connection
	databases.ConnectDatabase()

	// init engine
	app := gin.Default()

	// init route
	routes.InitRoute(app)

	// run app
	app.Run(app_config.PORT)
}
