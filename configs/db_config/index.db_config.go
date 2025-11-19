package db_config

import "os"

var DB_DRIVER = "mysql"
var DB_HOST = "127.0.0.1"
var DB_PORT = "3306"
var DB_NAME = "gogin"
var DB_USERNAME = "root"
var DB_PASSWORD = ""

func InitDBConfig() {
	driverEnv := os.Getenv("DB_DRIVER")
	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	hostEnv := os.Getenv("DB_HOST")
	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")
	if portEnv != "" {
		DB_PORT = portEnv
	}

	nameEnv := os.Getenv("DB_NAME")
	if nameEnv != "" {
		DB_NAME = nameEnv
	}

	usernameEnv := os.Getenv("DB_USERNAME")
	if usernameEnv != "" {
		DB_USERNAME = usernameEnv

	}
	passwordEnv := os.Getenv("DB_PASSWORD")
	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}
}
