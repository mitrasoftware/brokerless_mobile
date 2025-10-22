package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model

	Id                 uint           `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	ProductID          string         `gorm:"column:product_id;not null" json:"product_id"`
	CategoryIcon       uint           `gorm:"column:category_icon;not null" json:"category_icon"`
	ProductName        string         `gorm:"column:product_name;not null" json:"product_name"`
	ShopId             string         `gorm:"column:shop_id;not null" json:"shop_id"`
	SellingPrice       string         `gorm:"column:selling_price;not null" json:"selling_price"`
	Mrp                string         `gorm:"column:mrp;not null" json:"mrp"`
	SearchKey          string         `gorm:"column:search_key" json:"search_key"`
	ImageUrl           string         `gorm:"column:image_url;not null" json:"image_url"`
	ProductDescription string         `gorm:"column:description;not null" json:"description"`
	Specifications     string         `gorm:"column:specifications" json:"specifications"`
	PurchasedPrice     string         `gorm:"column:purchase_price;not null" json:"purchase_price"`
	DeliveryCharge     string         `gorm:"column:delivery_charge" json:"delivery_charge"`
	AvailableQuantity  string         `gorm:"column:quantity;not null" json:"quantity"`
	Blurhash           string         `gorm:"column:blurhash" json:"blurhash"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type ProductsRequest struct {
	Id                 uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	ProductID          string `gorm:"column:product_id;not null" json:"product_id"`
	CategoryIcon       uint   `gorm:"column:category_icon;not null" json:"category_icon"`
	ProductName        string `gorm:"column:product_name;not null" json:"product_name"`
	ShopId             string `gorm:"column:shop_id;not null" json:"shop_id"`
	SellingPrice       string `gorm:"column:selling_price;not null" json:"selling_price"`
	Mrp                string `gorm:"column:mrp;not null" json:"mrp"`
	SearchKey          string `gorm:"column:search_key" json:"search_key"`
	ImageUrl           string `gorm:"column:image_url;not null" json:"image_url"`
	ProductDescription string `gorm:"column:description;not null" json:"description"`
	Specifications     string `gorm:"column:specifications" json:"specifications"`
	PurchasedPrice     string `gorm:"column:purchase_price;not null" json:"purchase_price"`
	DeliveryCharge     string `gorm:"column:delivery_charge" json:"delivery_charge"`
	AvailableQuantity  string `gorm:"column:quantity;not null" json:"quantity"`
	Blurhash           string `gorm:"column:blurhash" json:"blurhash"`
}
