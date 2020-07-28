package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"github.com/jasongauvin/wikiPattern/services"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetRegistrationForm(c *gin.Context)  {
	c.HTML(
		http.StatusOK,
		"security/registration.html",
		gin.H{
			"title":    "Registration",
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
	userCreated, err := services.SaveUser(registerForm.Email, registerForm.Password)
	if err != nil {
		c.HTML(
			http.StatusUnprocessableEntity,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	sessionKey := uuid.NewV4().String()
	session, err := models.CreateUserSession(userCreated, sessionKey)
	if err != nil {
		c.HTML(
			http.StatusUnprocessableEntity,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("session_token", session.SessionKey, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusMovedPermanently, "/auth/profile")
	c.Abort()
}

func GetLoginForm(c *gin.Context)  {
	c.HTML(
		http.StatusOK,
		"security/login.html",
		gin.H{
			"title":    "Registration",
		})
}

func Login(c *gin.Context)  {
	var err error
	var loginForm services.LoginForm
	if err = c.ShouldBind(&loginForm); err != nil {
		fmt.Println("error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	var userSession *models.UserSession
	userSession, err = services.AuthenticateUser(loginForm.Email, loginForm.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if userSession != nil {
		c.SetCookie("session_token", userSession.SessionKey , 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusMovedPermanently, "/auth/profile")
		c.Abort()
	}
}
