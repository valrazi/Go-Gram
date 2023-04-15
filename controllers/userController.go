package controllers

import (
	"go-final-task/database"
	"go-final-task/helpers"
	"go-final-task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

// RegisterUser	godoc
// @Summary Register new user
// @Description Response with user that recently added
// @Tags user
// @Produce json
// @Success 200 {object} models.User
// @Router /users/register [post]
// @Param socialMedia body models.UserRegisterDTO true "User JSON"
func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

// LoginUser	godoc
// @Summary Login to get token
// @Description Response with token that
// @Tags user
// @Produce json
// @Success 200 {object} models.TokenDTO
// @Router /users/login [post]
// @Param socialMedia body models.UserLoginDTO true "User JSON"
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	_ = db

	User := models.User{}
	password := ""

	c.ShouldBindJSON(&User)
	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.CompareHash([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
