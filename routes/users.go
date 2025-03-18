package routes

import (
	"event-booking/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.ID = 1
	user.Save()
	context.JSON(201, gin.H{"message": "User created successfully", "user": user})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = user.Login()

	if err != nil {
		context.JSON(401, gin.H{"message": "Invalid credentials", "error": err})
		return
	}

	context.JSON(200, gin.H{"message": "Login successful"})
}
