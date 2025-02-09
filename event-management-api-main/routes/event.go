package routes

import (
	"project/event-management-api/controllers"
	"project/event-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func EventRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", middlewares.Authenticate, controllers.CreateEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
	authenticated.POST("/events/:id/register", controllers.RegisterEvent)
	authenticated.DELETE("/events/:id/register", controllers.CancelRegistrationEvent)
}
