package handler

import (
	"net/http"

	"shoppers/internal/dto"
	"shoppers/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	userID := c.GetString("user_id")

	var req dto.CreateAddressRequest

	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})

		return
	}

	err = service.CreateAddress(
		userID,
		&req,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message" : "address create successfully",
	})
}

func GetAddressByUserID(c *gin.Context) {
	userID := c.GetString("user_id")

	address, err := service.GetAddressByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, address)
}