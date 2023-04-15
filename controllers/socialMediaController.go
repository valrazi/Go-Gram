package controllers

import (
	"go-final-task/database"
	"go-final-task/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// PostSocialMedia	godoc
// @Summary Add new social media
// @Description Response with social media that recently added
// @Tags social-media
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /social-media [post]
// @Param socialMedia body models.SocialMediaDTO true "Social Media JSON"
// @Param Authorization header string true "Insert your token here" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
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

// GetAllSocialMedia	godoc
// @Summary Get all social media
// @Description Response with all social media that exists
// @Tags social-media
// @Produce json
// @Success 200 {array} models.SocialMedia
// @Router /social-media [get]
// @Param Authorization header string true "Insert your token here" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
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

// GetSocialMediaByID	godoc
// @Summary Get social media by ID
// @Description Response specified social media by ID as JSON
// @Tags social-media
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /social-media/{smID} [get]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param smID path string true  "Search by ID"
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

// DeleteSocialMediaByID	godoc
// @Summary Delete social media by ID
// @Description Response with HTTP Status
// @Tags social-media
// @Produce json
// @Success 200
// @Router /social-media/{smID} [delete]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param smID path string true "Delete by ID"
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
		"message": "Succesfully deleted social media",
	})
}

// UpdateSocialMediaByID	godoc
// @Summary Update social-media by ID
// @Description Response with updated social-media
// @Tags social-media
// @Produce json
// @Success 200 {object} models.SocialMedia
// @Router /social-media/{smID} [put]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param smID path string true "Update by ID"
// @Param social-media body models.SocialMediaDTO true "Comment JSON"
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
