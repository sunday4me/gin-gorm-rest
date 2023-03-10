package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sunday4me/gin-gorm-rest/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)

}
