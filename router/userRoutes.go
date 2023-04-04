package router

import (
	"grpc-auth-micro/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.Engine){

	c.GET("/",controllers.UserHome)
	c.POST("/login",controllers.Login)
	c.POST("/signup",controllers.Signup)

}