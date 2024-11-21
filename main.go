package main

import (
	"rest-api/db"
	"rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB() // Initialize the database connection
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

