package controllers

import (
	"go-final-task/database"
	"go-final-task/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// PostPhoto godoc
// @Summary Add new photo
// @Description Response with photo that recently added
// @Tags photo
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photo [post]
// @Param comment body models.PhotoDTO true "Comment JSON"
// @Param Authorization header string true "Insert your token here" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
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

// GetAllPhoto	godoc
// @Summary Get all photo
// @Description Response with all photo that exists
// @Tags photo
// @Produce json
// @Success 200 {array} models.Photo
// @Router /photo [get]
// @Param Authorization header string true "Insert your token here" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
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

// GetPhotoByID	godoc
// @Summary Get photo by ID
// @Description Response specified photo by ID as JSON
// @Tags photo
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photo/{photoID} [get]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param photoID path string true  "Search by ID"
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

// DeletePhotoByID	godoc
// @Summary Delete photo by ID
// @Description Response with HTTP Status
// @Tags photo
// @Produce json
// @Success 200
// @Router /photo/{photoID} [delete]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param photoID path string true "Delete by ID"
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

// UpdatePhotoByID	godoc
// @Summary Update photo by ID
// @Description Response with updated photo
// @Tags photo
// @Produce json
// @Success 200 {object} models.Photo
// @Router /photo/{photoID} [put]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param photoID path string true "Update by ID"
// @Param photo body models.PhotoDTO true "Photo JSON"
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
