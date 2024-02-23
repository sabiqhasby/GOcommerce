package models

import "gorm.io/gorm"

type ProductCategory struct {
	gorm.Model
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
