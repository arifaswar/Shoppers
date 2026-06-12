package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoppers/internal/service"
)

func Checkout(c *gin.Context) {
	userID := c.GetString("user_id")

	response, err := service.Checkout(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}