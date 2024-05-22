package routes

import (
	"CloudSystem/controller"

	"github.com/gin-gonic/gin"
)

func RegisterFirewallRoutes(router *gin.Engine) {
	FirewallGroup := router.Group("/firewall")
	{
		FirewallGroup.POST("/", controller.AddFirewall)
		FirewallGroup.GET("/", controller.GetAllFirewalls)
	}
}
