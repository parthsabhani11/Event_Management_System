package routes

import (
	"project/event-management-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(server *gin.Engine) {
	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.Login)
}
