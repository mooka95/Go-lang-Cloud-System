package routes

import (
	"CloudSystem/controller"
	"CloudSystem/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterFirewallRoutes(router *gin.Engine) {
	FirewallGroup := router.Group("/firewall")
	{
		FirewallGroup.Use(middlewares.Authenticate)
		FirewallGroup.POST("/", controller.AddFirewall)
		FirewallGroup.GET("/", controller.GetAllFirewalls)
		FirewallGroup.GET("/:id", controller.GetFirewallByID)
		FirewallGroup.DELETE("/:id", controller.DeleteFirewall)
	}
}
