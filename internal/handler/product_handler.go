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

func GetProductById(c *gin.Context) {
	id := c.Query("id")

	product, err := service.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product" : product,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Query("id")

	var req CreateProductRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	err = service.UpdateProduct(id, req.Name, req.Description, req.Price, req.Stock, req.ImageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Product updated successfully",
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Query("id")

	err := service.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Product deleted successfully",
	})
}