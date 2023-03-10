package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sunday4me/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.GetUsers) 
	router.POST("/", controller.GetUsers) 
		

	}


