package controllers

import "github.com/gin-gonic/gin"

func ListServices(c *gin.Context) {

}

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"mime/multipart"
// 	"net/http"
// 	"sync"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/service/s3"
// 	"github.com/gin-gonic/gin"
// 	"github.com/mitrasoftware/pureone_backend_go/config"
// )

// func ListPGController(c *gin.Context) {
// 	blurhash := c.PostForm("blurhash")
// 	whoiam := c.PostForm("whoiam")
// 	name := c.PostForm("name")
// 	mobile := c.PostForm("mobile")
// 	propertyFor := c.PostForm("property_for")
// 	city := c.PostForm("city")
// 	subcity := c.PostForm("subcity")
// 	address := c.PostForm("address")
// 	pgName := c.PostForm("pg_name")
// 	pincode := c.PostForm("pincode")
// 	pgOperationSince := c.PostForm("pg_operation_since")
// 	landmark := c.PostForm("landmark")
// 	pgPresentIn := c.PostForm("pg_present_in")
// 	gender := c.PostForm("gender")
// 	tenantType := c.PostForm("tenant_type")
// 	roomFacilitiesListMap := c.PostForm("room_facilities_list_map")
// 	pgRulesListMap := c.PostForm("pg_rules_list_map")
// 	deviceID := c.PostForm("device_id")

// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
// 		return
// 	}

// 	files := form.File["image_url"]
// 	if len(files) == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No files uploaded"})
// 		return
// 	}

// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	imageUrls := make([]string, 0, len(files))
// 	errChan := make(chan error, len(files))

// 	for _, file := range files {
// 		wg.Add(1)
// 		go func(file *multipart.FileHeader) {
// 			defer wg.Done()

// 			f, openErr := file.Open()
// 			if openErr != nil {
// 				errChan <- fmt.Errorf("failed to open image: %w", openErr)
// 				return
// 			}
// 			defer f.Close()

// 			result, uploadErr := config.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
// 				Bucket: aws.String("broker-less1"),
// 				Key:    aws.String(file.Filename),
// 				Body:   f,
// 				ACL:    "public-read",
// 			})
// 			if uploadErr != nil {
// 				errChan <- fmt.Errorf("failed to upload image: %w", uploadErr)
// 				return
// 			}

// 			mu.Lock()
// 			imageUrls = append(imageUrls, result.Location)
// 			mu.Unlock()
// 		}(file)
// 	}

// 	wg.Wait()
// 	close(errChan)

// 	for err := range errChan {
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 	}

// 	jsonImages, marshalErr := json.Marshal(imageUrls)
// 	if marshalErr != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal image URLs"})
// 		return
// 	}

// 	query := "INSERT INTO pg_properties (id, created_at, updated_at, product_images, blurhash, whoiam, name, mobile, property_for, city, subcity, address, pg_name, pincode, pg_operation_since, landmark, pg_present_in, gender, tenant_type, room_facilities_list_map,pg_rules_list_map, device_id) VALUES (NULL, NOW(), NOW(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
// 	stmt, err := config.DB.Prepare(query)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	result, err := stmt.Exec(string(jsonImages), blurhash, whoiam, name, mobile, propertyFor, city, subcity, address, pgName, pincode, pgOperationSince, landmark, pgPresentIn, gender, tenantType, roomFacilitiesListMap, pgRulesListMap, deviceID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database execution error"})
// 		return
// 	}

// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch affected rows"})
// 		return
// 	}

// 	// fmt.Println("Successfully uploaded and added to database", affectedRows)
// 	c.JSON(http.StatusOK, gin.H{"status": "success"})

// }
