package handlers

import (
	"gocommerce/models"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup

		//Menambahkan satu goroutine ke waitgrup
		wg.Add(1)

		//start goroutine untuk melakukan operasi yg membutuhkan waktu lama
		go func() {
			defer wg.Done() //menampilkan notif bahwa goroutine sudah selesai
			db.Find(&products)
		}()

		//menunggu goroutine selesai
		wg.Wait()

		ctx.JSON(200, products)
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var product models.Product

		//ambil data di database yg pertama muncul, jika kosong maka tampilkan errors
		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
		ctx.JSON(200, product)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.Product

		//Jika input dari browser kosong, maka munculkan error
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{
				"message": "invalid input",
			})
			return
		}
		db.Create(&input)
		ctx.JSON(201, input)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var product models.Product

		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}

		var input models.Product
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		db.Model(&product).Updates(input)
		ctx.JSON(200, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var product models.Product

		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{
				"message": "Product not found",
			})
			return
		}
		db.Delete(&product)

		ctx.JSON(200, gin.H{
			"message": "Product deleted",
		})
	}
}
