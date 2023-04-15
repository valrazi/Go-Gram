package router

import (
	"go-final-task/controllers"
	"go-final-task/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photo")

	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.AddPhoto)
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:photoID", controllers.GetPhotoByID)
		photoRouter.PATCH("/:photoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")

	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.AddComment)
		commentRouter.GET("/", controllers.GetAllComent)
		commentRouter.GET("/:commentID", controllers.GetCommentById)
		commentRouter.PATCH("/:commentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/social-media")

	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.AddSocialMedia)
		socialMediaRouter.GET("/", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:smID", controllers.GetSocialMediaById)
		socialMediaRouter.PATCH("/:smID", controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:smID", controllers.DeleteSocialMedia)
	}
	return r
}
