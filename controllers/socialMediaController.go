package controllers

import (
	"go-final-task/database"
	"go-final-task/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AddSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	c.ShouldBindJSON(&SocialMedia)

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	SocialMedia := models.SocialMedia{}
	var socialMedia []models.SocialMedia
	userID := uint(userData["id"].(float64))

	SocialMedia.UserID = userID

	err := db.Debug().Find(&socialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, socialMedia)
}

func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))
	socialMediaID, _ := strconv.Atoi(c.Param("smID"))

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).First(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))
	socialMediaID, _ := strconv.Atoi(c.Param("smID"))

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Delete(&SocialMedia).Error

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

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))
	socialMediaID, _ := strconv.Atoi(c.Param("smID"))

	c.ShouldBindJSON(&SocialMedia)

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaURL: SocialMedia.SocialMediaURL,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SocialMedia)
}
