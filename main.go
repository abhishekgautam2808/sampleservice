package main

import (
	"github.com/abhishekgautam2808/sampleservice/controllers"
	"github.com/abhishekgautam2808/sampleservice/models"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	models.ConnectDatabase()

	server.GET("/user", controllers.FindAll)
	server.GET("/user/:id", controllers.GetSingleUser)
	server.Run(":8080")
}
