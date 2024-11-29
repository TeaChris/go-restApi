package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST(("/login"), login)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:eventId", getEvent)
	server.PUT("/events/:eventId", updateEvent)
	server.DELETE("/events/:eventId", deleteEvent)
}