package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	configaws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mitrasoftware/pureone_backend_go/config"
	"github.com/mitrasoftware/pureone_backend_go/models"

	"github.com/gin-gonic/gin"
)

// AddProducts godoc
// @Summary Add a new product
// @Description Adds a product under a category
// @Tags Product
// @Accept  multipart/form-data
// @Produce  json
// @Param category_id formData int true "Category ID"
// @Param title formData string true "Product title"
// @Param description formData string false "Product description"
// @Param image formData file false "Product image"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /api/add_products [post]

func AddProducts(c *gin.Context) {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var imageUrls []string

	categoryId := c.PostForm("category_id")
	productName := c.PostForm("product_name")
	shopId := c.PostForm("shop_id")
	sellingPrice := c.PostForm("selling_price")
	mrp := c.PostForm("mrp")
	searchKey := c.PostForm("search_key")
	description := c.PostForm("description")
	specification := c.PostForm("specifications")
	purchasedPrice := c.PostForm("purchase_price")
	deliveryCharges := c.PostForm("delivery_charges")
	quantity := c.PostForm("quantity")
	blurhash := c.PostForm("blurhash")

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

	//  AWS Config
	cfg, err := configaws.LoadDefaultConfig(context.TODO(),
		configaws.WithRegion("ap-south-1"),
	)
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	//  Loop through files and upload concurrently
	for _, file := range files {
		wg.Add(1)

		go func(f *multipart.FileHeader) {
			defer wg.Done()

			fileReader, err := f.Open()
			if err != nil {
				log.Println("Failed to open file:", err)
				return
			}
			defer fileReader.Close()

			//  Upload to S3
			_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
				Bucket: aws.String("pureone-storage"),
				Key:    aws.String(f.Filename),
				Body:   fileReader,

				//  ACL: aws.String("public-read"),
				//  Remove if bucket has BlockPublicACLs
			})
			if err != nil {
				log.Println("Failed to upload file:", err)
				return
			}

			//  Build URL
			imageURL := fmt.Sprintf("https:// %s.s3.%s.amazonaws.com/%s",
				"pureone-storage", "ap-south-1", f.Filename)

			//  Append URL safely
			mu.Lock()
			imageUrls = append(imageUrls, imageURL)
			mu.Unlock()

		}(file)
	}

	//  Wait for all uploads to finish
	wg.Wait()
	catID64, err := strconv.ParseUint(categoryId, 10, 32) // base 10, 32-bit size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		panic(err)
	}
	catID := uint(catID64) // convert to uint

	jsonBytes, err := json.Marshal(imageUrls)
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonBytes)

	product := models.Products{
		CategoryIcon:       catID,
		ProductName:        productName,
		ShopId:             shopId,
		SellingPrice:       sellingPrice,
		Mrp:                mrp,
		SearchKey:          searchKey,
		ProductDescription: description,
		Specifications:     specification,
		PurchasedPrice:     purchasedPrice,
		DeliveryCharge:     deliveryCharges,
		AvailableQuantity:  quantity,
		Blurhash:           blurhash,
		ImageUrl:           jsonString,
	}

	// Insert into database
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//  All images uploaded — now respond or insert into DB
	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
	})

}
