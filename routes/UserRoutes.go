package routes

import (
	"github.com/gin-gonic/gin"
	"go-rest-service/controller"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.PostUser)
	router.POST("/excel", controller.CreateExcel)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.PutUser)
	router.GET("/:id", controller.GetUserById)

}
