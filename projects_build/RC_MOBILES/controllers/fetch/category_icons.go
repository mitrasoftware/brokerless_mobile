package controllers

import (
	"github.com/gin-gonic/gin"
)

func CategoryIcons(c *gin.Context) {

	// var services []models.CategoryIcons

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	// query := "SELECT title, image_url FROM category_icons"

	// rows, err := config.DB.Query(query)

	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var service models.CategoryIcons
	// 	err := rows.Scan(&service.Title, &service.ImageURL)
	// 	if err != nil {
	// 		panic(err)

	// 	}
	// 	services = append(services, service)
	// }

	// c.JSON(http.StatusOK, services)

}
