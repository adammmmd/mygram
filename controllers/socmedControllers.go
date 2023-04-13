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


func GetAllSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Socialmedias := []models.SocialMedia{}

	err := db.Debug().Where("user_id = ?", userID).Find(&Socialmedias).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cant find the social media",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"socmed": Socialmedias,
	})
}

func GetOneSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Socialmedias := []models.SocialMedia{}

	socmedId, _ := strconv.Atoi(c.Param("socmedId"))
	err := db.Debug().Where("user_id = ?", userID).Where("id = ?", socmedId).Find(&Socialmedias).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cant find the photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"socmed": Socialmedias,
	})
}

func CreateSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socmed := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = userID

	err := db.Debug().Create(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"created": Socmed,
	})
}

func UpdateSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Socmed := models.SocialMedia{}

	socmedId, _ := strconv.Atoi(c.Param("socmedId"))
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = userID
	Socmed.ID = uint(socmedId)

	err := db.Model(&Socmed).Where("id = ?", socmedId).Updates(models.SocialMedia{Name: Socmed.Name, SocialMediaURL: Socmed.SocialMediaURL}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated": Socmed,
	})
}

func DeleteSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socmedId, _ := strconv.Atoi(c.Param("socmedId"))
	Socmed := models.SocialMedia{}

	err := db.Debug().Where("id = ?", socmedId).Where("user_id = ?", userID).First(&Socmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Delete(&Socmed).Error
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