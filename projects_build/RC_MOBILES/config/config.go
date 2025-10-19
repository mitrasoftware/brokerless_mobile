package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Uploader *manager.Uploader

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL DB from GORM: ", err)
	}

	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping PostgreSQL database: ", err)
	}

	DB = db
	log.Println("PostgreSQL database connected successfully")

	// Initialize AWS S3 uploader
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println("Error loading AWS config:", err)
	}

	client := s3.NewFromConfig(cfg)
	Uploader = manager.NewUploader(client)
	log.Println("AWS S3 Uploader initialized successfully")
}
