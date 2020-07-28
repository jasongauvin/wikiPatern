package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"net/http"
)

func GetProfilePage(c *gin.Context) {
	var userSession *models.UserSession
	var user *models.User
	var err error
	cookie, err := c.Cookie("session_token")
	if err != nil {
		c.HTML(
			http.StatusUnauthorized,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	userSession, err = models.FindSessionByKey(cookie)
	if err != nil {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	user, err = models.FindUserByID(userSession.UserId)
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
