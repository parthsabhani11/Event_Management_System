package controllers

import (
	"net/http"
	"project/event-management-api/models"
	"project/event-management-api/utils"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Wrong user input format. Please check the format and try again."})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Wrong user input format. Please check the format and try again."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
