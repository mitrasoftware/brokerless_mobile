package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mitrasoftware/pureone_backend_go/config"
	"github.com/mitrasoftware/pureone_backend_go/models"
)

// CategoryIcons godoc
// @Summary Get category icons
// @Description Fetch list of category icons (requires JWT token in Authorization header)
// @Tags Category
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string][]models.CategoryIconsResponse "List of category icons"
// @Failure 401 {object} map[string]string "Unauthorized – missing or invalid token"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/fetch_category_icons [get]
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
