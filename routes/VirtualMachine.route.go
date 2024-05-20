package routes

import (
    "github.com/gin-gonic/gin"
    "CloudSystem/controller"
)


func RegisterVirtualMachinesRoutes(router *gin.Engine) {
    virtualmachineGroup := router.Group("/virtualmachine")
    {
        virtualmachineGroup.POST("/",controller.AddVirtualMachine)
    }
}
