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

func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Comments := []models.Comment{}

	err := db.Debug().Where("user_id = ?", userID).Find(&Comments).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cant find the comments",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": Comments,
	})
}

func GetOneComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	err := db.Debug().Where("user_id = ?", userID).Where("id = ?", commentId).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "cant find the photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Comment": Comment,
	})
}

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"created": Comment,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == "application/json" {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{PhotoID: Comment.PhotoID, Message: Comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated": Comment,
	})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	err := db.Debug().Where("id = ?", commentId).Where("user_id = ?", userID).First(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Delete(&Comment).Error
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
