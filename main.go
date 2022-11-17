package main

import (
	"asignment/controllers"
	"asignment/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	router := gin.Default()
	router.POST("/createUser", controllers.AddUsers)
	router.Run(":8080")
}
