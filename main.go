package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sunday4me/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()

	//router.GET("/", func(c *gin.Context){
	//c.String(200, "Hello World!")
	// })
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}
