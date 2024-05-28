package controller

import (
	"CloudSystem/database"
	"CloudSystem/models"
	"CloudSystem/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	//start transaction
	//check if email exists
	currentConnection, err := database.DB.Begin()
	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	}
	defer func() {
		if err != nil {
			currentConnection.Rollback()
		}
	}()

	var user models.User
	err = context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	// user.Id = context.GetInt64("userId")
	_, err = models.GetUserByEmail(user.Email)
	if err == nil {
		context.JSON(http.StatusConflict, gin.H{"message": "email already exists"})
		return
	}
	_, err = user.AddUser(currentConnection)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}
	_, err = user.AddUserAddress(currentConnection)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later."})
		return
	}
	err = currentConnection.Commit()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created Successfully", "userId": user.Identifier})
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
