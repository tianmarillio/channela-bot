package main

import (
	"github.com/tianmarillio/channela-backend/config"
	"github.com/tianmarillio/channela-backend/src/models"
)

func main() {
	config.LoadEnv()
	config.ConnectToDB()
	config.DB.AutoMigrate(
		&models.Channel{},
		&models.User{},
	)
}
