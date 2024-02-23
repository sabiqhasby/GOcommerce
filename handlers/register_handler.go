package handlers

import (
	"gocommerce/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.User
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid Input"})
			return
		}

		//cek username is already taken
		var existingUser models.User
		if err := db.Where("username = ?", input.Username).First(&existingUser).Error; err != nil {
			ctx.JSON(400, gin.H{"message": "Username already exists"})
			return
		}

		//Create new User

		newUser := models.User{
			Username: input.Username,
			Password: input.Password, // find hashing later, and make sure it save before storing it to db
		}

		if err := db.Create(&newUser).Error; err != nil {
			ctx.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

		//Optional, generate token for new user after registration

		token, err := CreateToken(newUser.ID)
		if err != nil {
			ctx.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "user registered successfully", "token": token,
		})
	}
}
