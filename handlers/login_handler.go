package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.User
		if err := ctx.ShouldBind(&input); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Input"})
			return
		}

		var user models.User
		if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
			ctx.JSON(401, gin.H{"message": "Invalid Credentials"})
			return
		}

		//pastikan memeriksa kata sandi yg benar disini

		token, err := CreateToken(user.ID)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

		ctx.JSON(200, gin.H{
			"token": token,
		})
	}
}
