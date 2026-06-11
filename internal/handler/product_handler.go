package handler

import (
	"net/http"
	"shoppers/internal/service"
	
	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock" binding:"required"`
	ImageURL    string  `json:"image_url"`
}

func CreateProduct(c *gin.Context) {

	var req CreateProductRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	err = service.CreateProduct(
		req.Name,
		req.Description,
		req.Price,
		req.Stock,
		req.ImageURL,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message" : "Product created successfully",
	})
}

func GetProducts(c *gin.Context) {

	products, err := service.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products" : products,
	})
}