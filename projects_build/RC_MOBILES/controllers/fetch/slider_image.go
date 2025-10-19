package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetSliderImages(c *gin.Context) {
	// var sliderArray []models.Slider

	// query := "SELECT id, image_url, categories, title, blurhash FROM slider_images"
	// rows, err := config.DB.Query(query)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var slider models.SliderImage
	// 	err := rows.Scan(&slider.ID, &slider.ImageURL, &slider.Categories, &slider.Title, &slider.Blurhash)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	sliderArray = append(sliderArray, slider)
	// }

	// if err = rows.Err(); err != nil {
	// 	panic(err)
	// }

	// if sliderArray != nil {
	// 	c.JSON(http.StatusOK, gin.H{

	// 		"slider": sliderArray,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{

	// 		"slider": []interface{}{},
	// 	})
	// }
}
