package main

import (
	"shoppers/internal/database"
	"shoppers/internal/domain"
	"shoppers/internal/config")

func main() {
	config.LoadEnv()

	database.ConnectDatabase()

	err := database.DB.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		panic(err)
	}
}