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
	firewall.UserId = context.GetInt64("userId")
	firewallIdentifier, err := firewall.InsertFirewall()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Firewall. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf(" Firewall Added successfully with id: %s", *firewallIdentifier)})
}
func GetAllFirewalls(context *gin.Context) {
	firewalls, err := models.GetAllFirewalls()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch firewalls. Try again later."})
		return
	}
	context.JSON(http.StatusOK, firewalls)
}
func GetFirewallByID(context *gin.Context) {
	// Extract the path parameter
	id := context.Param("id")
	firewall, err := models.GetFirewallByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Firewall. Try again later."})
		return
	}
	context.JSON(http.StatusOK, firewall)
}
func DeleteFirewall(context *gin.Context) {
	firewallId := context.Param("id")

	firewall, err := models.GetFirewallByID(firewallId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Firewall not exists"})
		return
	}
	err = firewall.DeleteFirewall()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the Firewall"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Firewall deleted successfully!"})

}
