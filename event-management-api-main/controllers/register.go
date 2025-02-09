package controllers

import (
	"net/http"
	"project/event-management-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user for event"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "User registered into event successfully"})
}

func CancelRegistrationEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse event ID"})
	}

	var event models.Event
	event.ID = eventId

	err = event.Unregister(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel registration for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User unregistered from event successfully"})
}
