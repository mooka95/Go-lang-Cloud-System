package controller

import (
	"CloudSystem/models"
	"CloudSystem/utils"
	"fmt"
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
	virtualMachine.UserId = context.GetInt64("userId")
	vmIdentifier, err := virtualMachine.InsertVirtualMachine()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create virtualMachine. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "virtualMachine created Successfully", "vmId": vmIdentifier})
}
func GetAllVirtualMachines(context *gin.Context) {
	virtualMachines, err := models.GetAllVirtualMachines()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch VirtualMachines. Try again later."})
		return
	}
	context.JSON(http.StatusOK, virtualMachines)
}
func GetVirtualMachineByID(context *gin.Context) {
	// Extract the path parameter
	id := context.Param("id")
	virtualMachines, err := models.GetVirtualMachineByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch VirtualMachine. Try again later."})
		return
	}
	context.JSON(http.StatusOK, virtualMachines)
}
func UpdateVirtualMachineActiveState(context *gin.Context) {
	body, err := utils.ExtractBodyFromRequest(context.Request.Body)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "SomeThing Went Wrong"})
		return
	}
	id, IdExists := body["id"]

	if !IdExists {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Missing id in the request body"})
		return
	}
	virtualMachine, err := models.GetVirtualMachineByID(id.(string))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "this virtual machine not exist "})
		return
	}
	if virtualMachine.IsActive == body["isActive"].(bool) && virtualMachine.IsActive {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Virtual machine is already active "})
		return
	} else if virtualMachine.IsActive == body["isActive"].(bool) && !virtualMachine.IsActive {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Virtual machine is already not Active "})
		return
	}
	err = virtualMachine.UpdateVirtualMachineActiveState(body["isActive"].(bool))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "can't Update VirtualMachine "})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf(" virtualMachine with id: %s updated successfully ", virtualMachine.Identifier)})
}
func DeleteVirtualMachine(context *gin.Context) {
	virtualMachineId := context.Param("id")

	virtualMachine, err := models.GetVirtualMachineByID(virtualMachineId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "VirtualMachine not exists"})
		return
	}
	err = virtualMachine.DeleteVirtualMachine(virtualMachine.Identifier)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the VirtualMachine"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "VirtualMachine deleted successfully!"})

}
