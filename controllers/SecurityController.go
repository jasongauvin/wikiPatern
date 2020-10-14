package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/services"
	"net/http"
)

func GetRegistrationForm(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"security/registration.html",
		gin.H{
			"title": "Registration",
		})
}

func Registration(c *gin.Context) {
	var registerForm services.RegisterForm
	if err := c.ShouldBind(&registerForm); err != nil {
		fmt.Println("error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	u, err := services.SaveUser(&registerForm)
	services.CreateUserSession(c, u)
	if err != nil {
		c.HTML(
			http.StatusUnprocessableEntity,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/auth/profile")
}

func GetLoginForm(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"security/login.html",
		gin.H{
			"title": "Registration",
		})
	return
}

func Login(c *gin.Context) {
	services.AuthenticateUser(c)
	return
}
