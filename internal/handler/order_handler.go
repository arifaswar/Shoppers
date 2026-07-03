package handler

import (
	"net/http"
	"shoppers/internal/dto"
	"shoppers/internal/service"

	"github.com/gin-gonic/gin"
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
		req.AddressID.String(),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetAllOrders(c *gin.Context) {

	orders, err := service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orders)
}
func GetMyOrders(c *gin.Context) {

	userID := c.GetString("user_id")
	orders, err := service.GetMyOrders(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c *gin.Context) {

	orderID := c.Param("id")
	order, err := service.GetOrderByID(orderID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrderStatus(c *gin.Context) {

	orderID := c.Param("id")

	var req dto.UpdateOrderStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	err := service.UpdateOrderStatus(
		orderID,
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

func CancelOrder(c *gin.Context) {

	orderID := c.Param("id")
	userID := c.GetString("user_id")
	err := service.CancelOrder(userID, orderID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order cancelled successfully",
	})
}
