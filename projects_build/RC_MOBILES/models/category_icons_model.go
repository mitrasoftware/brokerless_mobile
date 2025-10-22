package models

import (
	"time"

	"gorm.io/gorm"
)

type CategoryIcons struct {
	gorm.Model

	Id        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string         `gorm:"column:title" json:"title"`
	SubTitle  string         `gorm:"column:subtitle" json:"subtitle"`
	Category  string         `gorm:"column:category" json:"category"`
	Icon      string         `gorm:"column:icon" json:"icon"`
	BlurHash  string         `gorm:"column:blurhash" json:"blurhash"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type CategoryIconsResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	SubTitle string `json:"subtitle"`
	Icon     string `json:"icon"`
	Category string `json:"category"`
	BlurHash string `json:"blurhash"`
}
