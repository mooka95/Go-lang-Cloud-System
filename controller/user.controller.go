package controller

import (
	"CloudSystem/models"
	"CloudSystem/utils"
	"fmt"
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
func LoginUser(context *gin.Context){
	// extract body
body,err:= utils.ExtractBodyFromRequest(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "SomeThing Went Wrong"})
		return
	}
	//check if email or password 
email:=body["email"].(string)
password:=body["password"].(string)
	//get email and password 
user,err:=	models.GetUserByEmail(email)
if err != nil {
	fmt.Println(err)
	context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
	return
}
//validate credintials 
err =user.ValidatePassword(password)
if err != nil {
	context.JSON(http.StatusUnauthorized, gin.H{"message": "email or password invalid"})
	return
}
//generate token and get response
token, err:=utils.GenerateToken(user.Email,user.Identifier)
if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
	return
}

context.JSON(http.StatusOK, gin.H{"message": "Login successful!","token":token})
}
