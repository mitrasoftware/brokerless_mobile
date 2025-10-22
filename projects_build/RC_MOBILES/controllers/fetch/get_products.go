package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mitrasoftware/pureone_backend_go/config"
	"github.com/mitrasoftware/pureone_backend_go/models"
)

func GetProducts(c *gin.Context) {

	var limit string = c.DefaultQuery("limit", "5") // gives "1" if not provided
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var products []models.ProductsRequest

	if err := config.DB.Model(&models.Products{}).
		Select("id, product_id, category_icon, product_name, shop_id, selling_price, mrp, search_key, image_url, description, specifications, delivery_charge, quantity, blurhash").
		Find(&products).Limit(limitInt).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category icons"})

		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})

}
