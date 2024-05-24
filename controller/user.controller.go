package controller

import (
	"CloudSystem/models"
	"CloudSystem/utils"
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
	//add ddress

	context.JSON(http.StatusCreated, gin.H{"message": "user created Successfully", "userId": userIdentifier})
}
func LoginUser(context *gin.Context) {
	// extract body
	body, err := utils.ExtractBodyFromRequest(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "SomeThing Went Wrong"})
		return
	}
	//check if email or password
	email, isEmailExists := body["email"]
	password, isPasswordExists := body["password"]
	if !isEmailExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing email in the request body"})
		return
	}
	if !isPasswordExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing Password in the request body"})
		return
	}
	user, err := models.GetUserByEmail(email.(string))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
		return
	}
	err = user.ValidatePassword(password.(string))
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
		return
	}
	//generate token and get response
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
