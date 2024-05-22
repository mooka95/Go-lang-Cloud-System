package controller

import (
	"CloudSystem/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddFirewall(context *gin.Context) {
	var firewall models.Firewall
	err := context.ShouldBindJSON(&firewall)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	firewallIdentifier, err := firewall.InsertFirewall()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Firewall. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf(" Firewall Added successfully with id: %s", *firewallIdentifier)})
}
