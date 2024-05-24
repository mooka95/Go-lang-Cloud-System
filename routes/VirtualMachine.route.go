package routes

import (
    "github.com/gin-gonic/gin"
    "CloudSystem/controller"
    "CloudSystem/middlewares"
)


func RegisterVirtualMachinesRoutes(router *gin.Engine) {
    virtualmachineGroup := router.Group("/virtualmachines")
    {
        virtualmachineGroup.Use(middlewares.Authenticate)
        virtualmachineGroup.POST("/",controller.AddVirtualMachine)
        virtualmachineGroup.GET("/",controller.GetAllVirtualMachines)
        virtualmachineGroup.GET("/:id",controller.GetVirtualMachineByID)
        virtualmachineGroup.POST("/firewall/attach",controller.AttachVirtualMachineToFirewall)
        virtualmachineGroup.PATCH("/power",controller.UpdateVirtualMachineActiveState)
        virtualmachineGroup.DELETE("/:id",controller.DeleteVirtualMachine)
    }
}
