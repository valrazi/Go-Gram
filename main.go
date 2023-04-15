package main

import (
	"go-final-task/database"
	"go-final-task/router"

	_ "go-final-task/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title     My Gram API
// @description     Instagram 'clone' REST API using JWT auth.
// @host      localhost:8080
// @BasePath  /

func main() {
	database.StartDB()
	r := router.StartApp()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
