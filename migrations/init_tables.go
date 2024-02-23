package migrations

import (
	"gocommerce/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Product{},
		&models.ProductCategory{},
		&models.Transaction{},
		&models.TransactionItem{},
		&models.User{},
	)

	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})
}
