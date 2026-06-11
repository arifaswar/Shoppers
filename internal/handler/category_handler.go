package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shoppers/internal/service"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(c *gin.Context) {

	var req CreateCategoryRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	err = service.CreateCategory(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message" : "Category Created",
	})
}

func GetCategories(c *gin.Context) {
	categories, err := service.GetCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : categories,
	})
}

func GetCategoryById(c *gin.Context) {
	id := c.Query("id")

	category, err := service.GetCategoryById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : category,
	})
}