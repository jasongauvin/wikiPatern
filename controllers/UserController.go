package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/services"
)

func GetUserProfile(c *gin.Context) {
	services.GetProfilePage(c)
}
