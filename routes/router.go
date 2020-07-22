package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Poil au message",
		})
	})
}