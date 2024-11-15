/*
 * Filename: c:\Users\LENOVO\projects\GO\main.go
 * Path: c:\Users\LENOVO\projects\GO
 * Created Date: Wednesday, November 13th 2024, 8:35:52 am
 * Author: Boluwatife Olasunkanmi O.
 *
 * Copyright (c) 2024 PendulumHq
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	server.GET("/alive", getEvents)

	server.Run(":8080")
}

func getEvents(context *gin.Context){
	context.JSON(http.StatusOK, gin.H{
		"message": "Server is ALIVE!!!",
	})
}