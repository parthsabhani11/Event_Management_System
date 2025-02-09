package routes

import (
	"project/event-management-api/controllers"
	"project/event-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(server *gin.Engine) {
	server.GET("/users", middlewares.Authenticate, middlewares.IsAdmin, controllers.GetUsers)
	server.DELETE("/users/:id", middlewares.Authenticate, middlewares.IsAdmin, controllers.DeleteUser)
}
