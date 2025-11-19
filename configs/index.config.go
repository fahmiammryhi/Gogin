package configs

import (
	"Gogin/configs/app_config"
	"Gogin/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDBConfig()
}
