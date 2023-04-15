package controllers

import (
	"go-final-task/database"
	"go-final-task/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

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
		"message": "Succesfully deleted photo",
	})
}

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
