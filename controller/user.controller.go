package controller

import (
	"CloudSystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	userIdentifier, err := user.AddUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created Successfully", "userId": userIdentifier})
}
// func GetAllUsers(context *gin.Context) {
// 	// var user models.User
// 	// err := context.ShouldBindJSON(&user)

// 	// if err != nil {
// 	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
// 	// 	return
// 	// }
// 	userIdentifier, err := user.AddUser()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
// 		return
// 	}

// 	context.JSON(http.StatusCreated, gin.H{"message": "user created Successfully", "userId": userIdentifier})
// }
