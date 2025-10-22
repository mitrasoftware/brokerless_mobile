package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitrasoftware/pureone_backend_go/config"
	"github.com/mitrasoftware/pureone_backend_go/models"
)

func CategoryIcons(c *gin.Context) {

	var categoryIcons []models.CategoryIconsResponse

	if err := config.DB.Model(&models.CategoryIcons{}).
		Select("id, title, subtitle, icon, category, blurhash").
		Find(&categoryIcons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category icons"})

		return
	}
	c.JSON(http.StatusOK, gin.H{"data": categoryIcons})

}
