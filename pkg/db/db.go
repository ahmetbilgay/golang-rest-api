package db

import (
	"golang-rest-api/pkg/utils"
)

func ConnectDB() {
	dbUri := utils.LoadDotEnv("DB_URI")
	utils.ConnectDB(dbUri)
}
