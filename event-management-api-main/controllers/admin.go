package controllers

import (
	"net/http"
	"project/event-management-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	context.JSON(http.StatusOK, users)
}

func DeleteUser(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse ID"})
	}

	user, err := models.GetUserByID(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	err = user.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user with ID " + context.Param("id")})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "user": user})
}
