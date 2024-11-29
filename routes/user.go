package routes

import (
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := user.Save(); err != nil {
		if isUniqueConstraintError(err) {
			context.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	context.JSON(http.StatusCreated, gin.H{
	"message": "User created successfully",
	"user":    user,
})

}

func isUniqueConstraintError(err error) bool {
	return err.Error() == "UNIQUE constraint failed: users.email"
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}
