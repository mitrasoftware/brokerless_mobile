package models

import (
	"time"

	"gorm.io/gorm"
)

type CategoryIcons struct {
	gorm.Model

	Id        int            `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string         `gorm:"column:title" json:"title"`
	SubTitle  string         `gorm:"column:subtitle" json:"subtitle"`
	Icon      string         `gorm:"column:icon" json:"icon"`
	Category  string         `gorm:"column:category" json:"category"`
	BlurHash  string         `gorm:"column:blurhash" json:"blurhash"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
