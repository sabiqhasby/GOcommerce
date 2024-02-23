package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID     uint              `gorm:"primary_key" json:"id"`
	UserID uint              `json:"user_id"`
	Amount float64           `json:"amount"`
	Items  []TransactionItem `gorm:"foreignKey:TransactionID" json:"items"`
	Price  float64           `json:"price"`
}
