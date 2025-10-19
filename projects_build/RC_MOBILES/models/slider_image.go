package models

import (
	"time"

	"gorm.io/gorm"
)

type Slider struct {
	gorm.Model

	Id         int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string         `gorm:"column:title" json:"title"`
	Categories string         `gorm:"column:categories" json:"categories"`
	ImageUrl   string         `gorm:"column:image_url" json:"image_url"`
	BlurHash   string         `gorm:"column:blurhash" json:"blurhash"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
