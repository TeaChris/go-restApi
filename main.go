package main

import (
	"net/http"
	"rest-api/db"
	"rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB() // Initialize the database connection
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return
}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64) 

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventId)	

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var newEvent models.Event
	err := context.ShouldBindJSON(&newEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent.ID = 1 
	newEvent.UserId = 1

	err = newEvent.Save()
 
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": newEvent}) 
}