package controllers

import (
	"go-final-task/database"
	"go-final-task/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AddPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	c.ShouldBindJSON(&Photo)

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	var photos []models.Photo
	userID := uint(userData["id"].(float64))

	Photo.UserID = userID

	err := db.Debug().Find(&photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, photos)
}

func GetPhotoByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	photoID, _ := strconv.Atoi(c.Param("photoID"))

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err := db.Model(&Photo).Where("id = ?", photoID).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	photoID, _ := strconv.Atoi(c.Param("photoID"))

	c.ShouldBindJSON(&Photo)

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err := db.Model(&Photo).Where("id = ?", photoID).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully deleted photo",
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))
	photoID, _ := strconv.Atoi(c.Param("photoID"))

	c.ShouldBindJSON(&Photo)

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err := db.Model(&Photo).Where("id = ?", photoID).Updates(models.Photo{
		Title:    Photo.Title,
		Caption:  Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}
