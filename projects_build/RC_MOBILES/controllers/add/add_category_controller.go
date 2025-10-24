package controllers

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"

	"github.com/mitrasoftware/pureone_backend_go/config"
	s3Config "github.com/mitrasoftware/pureone_backend_go/config"
	"github.com/mitrasoftware/pureone_backend_go/models"
)

func AddCategory(c *gin.Context) {
	title := c.PostForm("title")
	subTitle := c.PostForm("subtitle")
	icon := c.PostForm("icon")
	category := c.PostForm("category")
	blurHash := c.PostForm("blurhash")

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	files := form.File["image_url"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var uploadedURLs []string

	// Parallel upload for all files
	for _, file := range files {
		wg.Add(1)

		go func(fileHeader *multipart.FileHeader) {
			defer wg.Done()

			f, openErr := fileHeader.Open()
			if openErr != nil {
				fmt.Println("Failed to open file:", openErr)
				return
			}
			defer f.Close()

			// Upload to S3
			_, uploadErr := s3Config.Uploader.Upload(context.TODO(), &s3.PutObjectInput{Bucket: aws.String("pureone-storage"), Key: aws.String(file.Filename), Body: f, ACL: "public-read"})

			if uploadErr != nil {
				fmt.Println("Upload error:", uploadErr)
				return
			}

			// Build URL manually
			imageURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s",
				"broker-less1", "ap-south-1", fileHeader.Filename)

			mu.Lock()
			uploadedURLs = append(uploadedURLs, imageURL)
			mu.Unlock()
		}(file)
	}

	wg.Wait()

	// Save record in DB
	categoryIcons := models.CategoryIcons{
		Title:    title,
		SubTitle: subTitle,
		Icon:     icon,
		Category: category,
		BlurHash: blurHash,
	}

	if err := config.DB.Create(&categoryIcons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"data":       categoryIcons,
		"image_urls": uploadedURLs,
	})
}
