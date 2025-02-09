package main

import (
	"net/http"
	"project/event-management-api/db"
	"project/event-management-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/", getDefault)

	routes.EventRoutes(server)
	routes.UserRoutes(server)
	routes.AdminRoutes(server)

	server.Run(":8080")
}

func getDefault(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World"}) //(200, "Hello World")
}
