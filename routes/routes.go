package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes() *gin.Engine{
	router := gin.Default()
	RegisterVirtualMachinesRoutes(router)

	return router
}