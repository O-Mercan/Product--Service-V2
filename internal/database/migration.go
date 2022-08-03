package database

import (
	"github.com/O-Mercan/Product--Service-V2/internal/product"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&product.Product{}); err != nil {
		return err
	}
	return nil
}
