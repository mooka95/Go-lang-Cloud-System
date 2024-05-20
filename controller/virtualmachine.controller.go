package controller

import (
	"CloudSystem/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddVirtualMachine(context *gin.Context) {
	var virtualMachine models.VirtualMachine
	err := context.ShouldBindJSON(&virtualMachine)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	vmIdentifier, err := virtualMachine.InsertVirtualMachine()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create virtualMachine. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "virtualMachine created Successfully", "vmId": vmIdentifier})
}
