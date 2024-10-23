package main

import (
	"github.com/gin-gonic/gin"
	"walnut/internal/service/ipml"
	"walnut/internal/web"
)

func main() {

	server := gin.Default()
	root := server.Group("/api").Group("/v1")
	userSrv := ipml.NewUserSrvImpl()
	userHandler := web.NewUserHandler(userSrv)
	userHandler.RegisterRoutes(root.Group("/user"))

	server.Run(":8080")
}
