package main

import (
	"go-final-task/database"
	"go-final-task/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
