package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTransactionWithItems(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var transaction models.Transaction
		if err := db.Preload("Items").First(&transaction, id).Error; err != nil {
			ctx.JSON(404, gin.H{
				"message": "Transaction not found",
			})
			return
		}
		ctx.JSON(200, transaction)
	}
}

func CreateTransaction(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.Transaction
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(404, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		// Jika Anda ingin mengaitkan item transaksi, pastikan item-item tersebut ada dalam database
		for _, item := range input.Items {
			var product models.Product
			if err := db.First(&product, item.ProductID).Error; err != nil {
				ctx.JSON(400, gin.H{"message": "Invalid Product ID"})
				return
			}
		}

		db.Create(&input)
		ctx.JSON(201, &input)
	}
}
