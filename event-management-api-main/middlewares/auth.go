package middlewares

import (
	"fmt"
	"net/http"
	"project/event-management-api/models"
	"project/event-management-api/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Please login first"})
		return
	}

	userId, err := utils.ValidateToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()
}

func IsAdmin(context *gin.Context) {
	userId := context.GetInt64("userId")

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	if user.Role != "admin" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("User %v is not an admin", user.Email)})
	}

	context.Next()
}
