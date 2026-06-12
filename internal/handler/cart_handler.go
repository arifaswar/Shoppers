package handler

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"shoppers/internal/dto"
	"shoppers/internal/service"
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

func GetCartItems(c *gin.Context) {

	userID := c.GetString("user_id")

	fmt.Printf("USER ID: %s\n", userID)

	cart, err := service.GetCartItems(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, cart)
}

func GetCartByUserId(c *gin.Context) {

	userID := c.GetString("user_id")
	cart, err := service.GetCartByUserId(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
			})
			return
	}
	c.JSON(http.StatusOK, cart)
}

func UpdateCartItem(c *gin.Context) {

	itemId := c.Query("id")
	if itemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "item_id query parameter is required",
		})
		return
	}

	itemID, err := uuid.Parse(itemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var req dto.UpdateCartItemRequest

	err = c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : err.Error(),
		})
		return
	}

	userID := c.GetString("user_id")

	err = service.UpdateCartItem(
		userID,
		itemID.String(),
		req.Quantity,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Cart item updated successfully",
	})
}

func DeleteCartItem(c *gin.Context) {

	itemId := c.Query("id")
	itemID, err := uuid.Parse(itemId)
	if itemId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message" : "item_id query parameter is required",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	userID := c.GetString("user_id")
	
	err = service.DeleteCartItem(userID, itemID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Cart item deleted successfully",
	})
}