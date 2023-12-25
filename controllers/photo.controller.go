package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/app"
	db "github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/database"
	"github.com/ranggakrisnaa/task-5-pbi-btpns-rangga/models"
)

func UploadPhotoProfile(ctx *gin.Context) {
	var reqBody app.PhotoReqBody
	userID, _ := ctx.Get("userID")

	file, err := ctx.FormFile("photo_file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing photo file"})
		return
	}

	timestamp := time.Now().UnixNano()
	filePhoto := fmt.Sprintf("%d-%s", timestamp, strings.ReplaceAll(file.Filename, " ", ""))
	filename := "http://localhost:3001/api/v1/photos/" + filePhoto

	reqBody.Title = ctx.PostForm("title")
	reqBody.Caption = ctx.PostForm("caption")
	reqBody.PhotoUrl = filename

	photo := models.Photo{
		Title:    reqBody.Title,
		Caption:  reqBody.Caption,
		PhotoUrl: reqBody.PhotoUrl,
		UserID:   userID.(string),
	}

	db := db.Init()

	err = ctx.SaveUploadedFile(file, "./uploads/"+filePhoto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	if err := db.Create(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Photo uploaded successfully",
		"success": true,
	})
}

func GetPhotoProfile(ctx *gin.Context) {

}

func UpdatePhotoProfile(ctx *gin.Context) {

}

func DeletePhotoProfile(ctx *gin.Context) {

}
