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

	user.Save()
	context.JSON(201, gin.H{"message": "User created successfully", "user": user})
}
