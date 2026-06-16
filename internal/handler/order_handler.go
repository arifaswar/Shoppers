package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoppers/internal/service"
	"shoppers/internal/dto"
)

func Checkout(c *gin.Context) {
	userID := c.GetString("user_id")

	var req dto.CheckoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := service.Checkout(
		userID,
		req.AddressID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetMyOrders(c *gin.Context) {

	userID := c.GetString("user_id")
	orders, err := service.GetMyOrders(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {

	orderID := c.Query("id")
	order, err := service.GetOrderByID(orderID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrderStatus(c *gin.Context) {

	id := c.Param("id")

	var req dto.UpdateOrderStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := service.UpdateOrderStatus(
		id,
		req.Status,
	)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order updated successfully",
	})
}