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
			protected.GET("/profile", handler.Profile)
		}

		product := api.Group("/products")
		{
			product.GET("", handler.GetProducts)
			product.GET(":id", handler.GetProductById)
			product.POST("", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.CreateProduct)
			product.PUT(":id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.UpdateProduct)
			product.DELETE(":id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.DeleteProduct)
		}

		category := api.Group("/category")
		{
			category.GET("", handler.GetCategories)
			category.POST("", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.CreateCategory)
			category.GET(":id", handler.GetCategoryById)
		}

		cart := api.Group("/cart")
		cart.Use(middleware.AuthMiddleware())
		{
			cart.POST("/add", handler.AddToCart)
			cart.GET("", handler.GetCartByUserId)
			cart.GET("/items", handler.GetCartItems)
			cart.PUT("/items", handler.UpdateCartItem)
			cart.DELETE("/items", handler.DeleteCartItem)
			cart.POST("/checkout", handler.Checkout)
		}

		order := api.Group("/orders")
		order.Use(middleware.AuthMiddleware())
		{
			order.GET("/my_orders", handler.GetMyOrders)
			order.GET("/:id", handler.GetOrderByID)
			order.PATCH("/:id/cancel", handler.CancelOrder)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			admin.GET("/orders", handler.GetAllOrders)
			admin.GET("/orders/:id", handler.GetOrderByID)
			admin.PATCH("/orders/:id/status", handler.UpdateOrderStatus)
		}

		address := api.Group("/address")
		address.Use(middleware.AuthMiddleware())
		{
			address.POST("", handler.CreateAddress)
			address.GET("", handler.GetAddressByUserID)
		}

		payment := api.Group("/payments")
		payment.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			payment.POST("/", handler.CreatePayment)
		}
	}
}
