package routes

import (
	"shoppers/internal/handler"
	"shoppers/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
		}

		protected := api.Group("/user")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/profile",handler.Profile)
		}

		product := api.Group("/products")
		{
			product.GET("/", handler.GetProducts)
			product.GET("/:id", handler.GetProductById)
			product.POST("/", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.CreateProduct)
			product.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.UpdateProduct)
			product.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.DeleteProduct)
		}
	}
}