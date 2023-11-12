package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-service/config"
	"go-rest-service/routes"
)

func main() {
	router := gin.Default()
	config.Connect()
	routes.UserRoutes(router)
	router.Run(":8084")
}
