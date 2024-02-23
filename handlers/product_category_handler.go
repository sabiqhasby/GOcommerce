package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET
func ListProductCategories(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categories []models.ProductCategory
		db.Find(&categories)
		ctx.JSON(200, categories)
	}
}

// GET BY ID
func GetProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var category models.ProductCategory
		if err := db.First(&category, id).Error; err != nil {

			ctx.JSON(404, gin.H{
				"message": "category Not found",
			})

			return
		}
		ctx.JSON(200, category)
	}
}

// POST, Create
func CreateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.ProductCategory

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"message": "invalid input"})
			return
		}

		if input.Name == "" {
			ctx.JSON(400, gin.H{"message": "Name is Required"})
		}

		db.Create(&input)
		ctx.JSON(201, input)
	}
}

// UPDATE, PUT
func UpdateProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var input, category models.ProductCategory

		if err := db.First(&category, id).Error; err != nil {
			ctx.JSON(404, gin.H{
				"message": "Category not found",
			})
			return
		}

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&category).Updates(input)
		ctx.JSON(200, category)
	}
}

// DELETE
func DeleteProductCategory(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var category models.ProductCategory

		if err := db.First(&category, id).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "Category not found"})
			return
		}

		db.Delete(&category)
		ctx.JSON(200, gin.H{
			"message": "Category deleted",
		})

	}
}
