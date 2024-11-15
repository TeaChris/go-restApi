package main

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context){
	var newEvent models.Event
	err := context.ShouldBindJSON(&newEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newEvent.ID = 1 // for now
	newEvent.UserId = 1 // for now

	newEvent.Save()
	
	context.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "event": newEvent}) 
}