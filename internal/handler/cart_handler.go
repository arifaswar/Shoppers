package handler

import (
	"net/http"
	"shoppers/internal/service"
	"github.com/gin-gonic/gin"
)

type AddCartRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func AddToCart(c *gin.Context) {

	userID := c.GetString("user_id")

	var req AddCartRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	err = service.AddToCart(
		userID,
		req.ProductID,
		req.Quantity,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Item added to cart successfully",
	})
}

func GetCart(c *gin.Context) {

	userID := c.GetString("user_id")

	cart, err := service.GetCart(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, cart)
}