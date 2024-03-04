package routes

import (
	"event-booking/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		log.Fatal("Could not register for event:", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event"})

}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Cancel(userId)

	if err != nil {
		log.Fatal("Could not cancel registration:", err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration cancelled for event"})
}
