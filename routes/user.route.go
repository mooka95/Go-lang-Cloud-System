package routes

import (
	"CloudSystem/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	UserGroup := router.Group("/user")
	{
		UserGroup.POST("/", controller.RegisterUser)
		UserGroup.POST("/search",controller.SearchWords)
	}
	router.POST("/login",controller.LoginUser)
}
