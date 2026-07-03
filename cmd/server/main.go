package main

import (
	"shoppers/internal/config"
	"shoppers/internal/database"
	"shoppers/internal/domain"
	"shoppers/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	database.ConnectDatabase()

	err := database.DB.AutoMigrate(
		&domain.User{},
		&domain.Product{},
		&domain.Category{},
		&domain.Cart{},
		&domain.CartItem{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Address{},
		&domain.Checkout{},
	)
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":3000")
}