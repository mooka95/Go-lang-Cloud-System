package controller

import (
	"CloudSystem/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAddress(context *gin.Context) {
	var address models.Address
	err := context.ShouldBindJSON(&address)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	address.UserId = context.GetInt64("userId")
	addressIdentifier, err := address.CreateAddress()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Address. Try again later."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf(" Address Added successfully with id: %s", *addressIdentifier)})
}
