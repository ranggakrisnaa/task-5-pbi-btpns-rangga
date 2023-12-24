package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/app"
	db "github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/database"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/models"
	"gorm.io/gorm"
)

func IsEmailExists(email string) bool {
	var user models.User

	db := db.Init()
	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

func IsUsernameExists(username string) bool {
	var user models.User

	db := db.Init()
	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

func Login(ctx *gin.Context) {
	var input app.LoginUser

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func Register(ctx *gin.Context) {
	var input app.RegisterUser

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if len(input.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Password must be at least 6 characters",
		})

		return
	}

	if IsEmailExists(input.Email) {
		ctx.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Email already exist, must be unique",
		})

		return
	}

	if IsUsernameExists(input.Username) {
		ctx.JSON(http.StatusConflict, gin.H{
			"success": false,
			"message": "Username already exist, must be unique",
		})

		return
	}

	u := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	db := db.Init()
	if err := db.Create(&u).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User registered successfully",
	})
}
