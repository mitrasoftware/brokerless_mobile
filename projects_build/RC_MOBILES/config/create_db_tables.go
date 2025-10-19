package config

import (
	"github.com/mitrasoftware/pureone_backend_go/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Slider{}, &models.CategoryIcons{})

}
