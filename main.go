package main

import (
	"grpc-auth-micro/database"
	"grpc-auth-micro/initializer"
	"grpc-auth-micro/router"
	"github.com/gin-gonic/gin"
)


func init(){
	initializer.LoadEnv()
	database.ConnectToDb()
	database.SyncDB()
}


func main(){
	rout :=gin.Default()
	router.UserRoutes(rout)
	rout.Run(":9000")
}