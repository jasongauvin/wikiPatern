package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"net/http"
)

func GetProfilePage(c *gin.Context) {
	var user *models.User
	var err error
	var us interface{}
	us, err = CheckSessionExistence(c)
	if err != nil {
		c.HTML(
			http.StatusUnauthorized,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	user, err = models.FindUserByUuiD(fmt.Sprintf("%v", us))
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "user/user_profile.html", gin.H{
		"title": "User profile",
		"user":  user,
	})
}
