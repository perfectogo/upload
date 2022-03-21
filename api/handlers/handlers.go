package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	l "github.com/perfectogo/upload/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *OutHandler) UploadImg(ctx *gin.Context) {

	file, err := ctx.FormFile("file")

	// The file cannot be received.
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	fmt.Println(extension)
	// Generate random file name
	newFileName := uuid.New().String() + extension
	filePath := "images/" + newFileName
	// The file is received, so let's save it
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		h.Log.Error("failed to upload image", l.Error(err))
		return
	}
	if err = h.Service.Upload().UploadImg(filePath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file in Database",
		})
		h.Log.Error("failed to upload image path in databse", l.Error(err))
		return
	}

	// File saved successfully. Return proper result
	ctx.JSON(201, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
}

func (h *OutHandler) GetImages(ctx *gin.Context) {
	imgURLs, err := h.Service.Upload().GetImages()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "not found images",
		})
		h.Log.Error("failed to get image from databse", l.Error(err))
		return
	}
	ctx.JSON(200, imgURLs)
}
