package models

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	ID            uint    `gorm:"primary_key" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
}
