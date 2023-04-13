package controllers

import (
	"Project/database"
	"Project/helpers"
	"Project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photos := []models.Photo{}

	err := db.Debug().Where("user_id = ?", userID).Find(&Photos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cant find the photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": Photos,
	})
}

func GetOnePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Photos := []models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	err := db.Debug().Where("user_id = ?", userID).Where("id = ?", photoId).Find(&Photos).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cant find the photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"photos": Photos,
	})
}

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"created": Photo,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated": Photo,
	})
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	err := db.Debug().Where("id = ?", photoId).Where("user_id = ?", userID).First(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "failed to delete photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been succesfully deleted",
	})
}