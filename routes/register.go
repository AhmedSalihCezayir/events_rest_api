package routes

import (
	"net/http"
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not found event with given id"})
		return
	}

	err = event.Register(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Registered!"})
}

func cancelRegistration(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not found event with given id."})
		return
	}

	err = models.CancelRegistration(event.ID, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}

func getEventAttendees(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	_, err = models.GetEventById(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not found event with given id."})
		return
	}

	attendees, err := models.GetEventAttendees(eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "There was a problem when fetching event attendees."})
		return
	}

	ctx.JSON(http.StatusOK, attendees)
}

func getUserRegisteredEvents(ctx *gin.Context) {
	userID := ctx.GetInt64("userID")
	
	events, err := models.GetRegisterations(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "There was a problem when fetching user registrations."})
		return
	}

	ctx.JSON(http.StatusOK, events)
}