package controller

import (
	"CloudSystem/models"
	"encoding/json"
	"fmt"
	"io"
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
	      // Extract the path parameter
		  var data map[string]interface{}
		  body, err := io.ReadAll(context.Request.Body)
		  if err != nil {
            context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read request body"})
            return
        }
		//   var requestBody map[string]interface{}
		  if err := json.Unmarshal(body, &data); err != nil {
			  context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			  return
		  }
		        // Access specific fields from the map
				id := data["id"].(string) // JSON numbers are decoded to float64 by default
				// isActive, isActiveExists := data["is_active"].(bool)
		  fmt.Println(id)
    
		  virtualMachine, err := models.GetVirtualMachineByID("ssss")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "this virtual machine not exist "})
		return
	}
	context.JSON(http.StatusOK, virtualMachine)
}
func DeleteVirtualMachine(context *gin.Context){
	virtualMachineId := context.Param("id")

	virtualMachine,err :=models.GetVirtualMachineByID(virtualMachineId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "VirtualMachine not exists"})
		return
	}
	err=virtualMachine.DeleteVirtualMachine(virtualMachine.Identifier)
	if(err!=nil){
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the VirtualMachine"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "VirtualMachine deleted successfully!"})

}