package controllers

import (
	"go-final-task/database"
	"go-final-task/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// PostComment	godoc
// @Summary Add new comment
// @Description Response with comment that recently added
// @Tags comment
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comment [post]
// @Param comment body models.CommentPostDTO true "Comment JSON"
// @Param Authorization header string true "Insert your token here" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
func AddComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	c.ShouldBindJSON(&Comment)

	Comment.UserID = userID

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

// GetComments	godoc
// @Summary Get All comment
// @Description Response with list of all comment as JSON
// @Tags comment
// @Produce json
// @Success 200 {array} models.Comment
// @Router /comment [get]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
func GetAllComent(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}
	var comment []models.Comment
	userID := uint(userData["id"].(float64))

	Comment.UserID = userID

	err := db.Debug().Find(&comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// GetCommentByID	godoc
// @Summary Get comment by ID
// @Description Response specified comment by ID as JSON
// @Tags comment
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comment/{commentID} [get]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param commentID path string true  "Search by ID"
func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	commentID, _ := strconv.Atoi(c.Param("commentID"))

	Comment.UserID = userID
	Comment.ID = uint(commentID)

	err := db.Model(&Comment).Where("id = ?", commentID).First(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

// DeleteCommentByID	godoc
// @Summary Delete comment by ID
// @Description Response with HTTP Status
// @Tags comment
// @Produce json
// @Success 200
// @Router /comment/{commentID} [delete]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param commentID path string true  "Delete by ID"
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	commentID, _ := strconv.Atoi(c.Param("commentID"))

	c.ShouldBindJSON(&Comment)

	Comment.UserID = userID
	Comment.ID = uint(commentID)

	err := db.Model(&Comment).Where("id = ?", commentID).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully deleted comment",
	})
}

// UpdateCommentByID	godoc
// @Summary Update comment by ID
// @Description Response with updated comment
// @Tags comment
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comment/{commentID} [put]
// @Param Authorization header string true "Insert your token" default(Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNoYW5hMTNAZ21haWwuY29tIiwiaWQiOjI0fQ.5-dveXe8_UMto_lpbHk3Nj9xWRQg-4b2vEfyuy6PK1k)
// @Param commentID path string true  "Update by ID"
// @Param comment body models.CommentUpdateDTO true "Comment JSON"
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))
	commentID, _ := strconv.Atoi(c.Param("commentID"))

	c.ShouldBindJSON(&Comment)

	Comment.UserID = userID
	Comment.ID = uint(commentID)

	err := db.Model(&Comment).Where("id = ?", commentID).Updates(models.Comment{
		Message: Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}
