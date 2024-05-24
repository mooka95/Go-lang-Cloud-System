package routes

import (
	"CloudSystem/controller"
	"CloudSystem/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterAddressRoutes(router *gin.Engine) {
	AddressGroup := router.Group("/address")
	{
		AddressGroup.Use(middlewares.Authenticate)
		AddressGroup.POST("/", controller.AddAddress)
		// AddressGroup.GET("/", controller.GetAllFirewalls)
		// AddressGroup.GET("/:id", controller.GetFirewallByID)
		// AddressGroup.DELETE("/:id", controller.DeleteFirewall)
	}
}
