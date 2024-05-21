package routes

import (
    "github.com/gin-gonic/gin"
    "CloudSystem/controller"
)


func RegisterVirtualMachinesRoutes(router *gin.Engine) {
    virtualmachineGroup := router.Group("/virtualmachines")
    {
        virtualmachineGroup.POST("/",controller.AddVirtualMachine)
        virtualmachineGroup.GET("/",controller.GetAllVirtualMachines)
        virtualmachineGroup.GET("/:id",controller.GetVirtualMachineByID)
    }
}
