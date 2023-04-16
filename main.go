package main

import (
	"go-final-task/database"
	"go-final-task/router"
	"os"

	_ "go-final-task/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     My Gram API
// @description     Instagram 'clone' REST API using JWT auth.
// @host      /
// @BasePath  /

func main() {
	database.StartDB()
	r := router.StartApp()

	var PORT = os.Getenv("PORT")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + PORT)
}
