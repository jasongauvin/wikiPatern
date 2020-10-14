package services

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/wikiPattern/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterForm struct {
	Name 	 string `form:"userName", binding:"required"`
	Email    string `form:"userEmail" binding:"required"`
	Password string `form:"userPassword" binding:"required"`
}

type LoginForm struct {
	Email    string `form:"userEmail" binding:"required"`
	Password string `form:"userPassword" binding:"required"`
}

// SaveUser creates a user.
// Return the created object
func SaveUser(registerForm *RegisterForm) (*models.User, error) {
	var err error
	user := models.User{
		Name:     registerForm.Name,
		Password: registerForm.Password,
		Email:    registerForm.Email,
		Admin:    false,
	}
	userCreated, err := models.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}

// AuthenticateUser let you find the user from its email,
// Compare the passed password and the hashed password stored.
// Check the session expiration date and update the token with a new one.
// Return the UserSession.
func AuthenticateUser(c *gin.Context) {
	var err error
	var user *models.User
	var loginForm LoginForm
	if err = c.ShouldBind(&loginForm); err != nil {
		fmt.Println("error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	user, err = models.FindUserByEmail(loginForm.Email)
	if err != nil {
		fmt.Println("error:", err)
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	err = models.VerifyPassword(user.Password, loginForm.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	CreateUserSession(c, user)
	c.Redirect(http.StatusMovedPermanently, "/auth/profile")
	return
}

func CreateUserSession(c *gin.Context, u *models.User)  {
	var err error
	session := sessions.Default(c)
	session.Set("uuid", u.Uuid)
	if err = session.Save(); err != nil {
		c.HTML(
			http.StatusBadRequest,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
}

// CheckSessionExistence let you find the userId by the sessionKey
// Return an error
func CheckSessionExistence(c *gin.Context) (interface{}, error) {
	session := sessions.Default(c)
	us := session.Get("uuid")
	if us == nil {
		return nil, errors.New("No session found")
	}

	return us, nil
}

// Logout let you delete the user's session
// Calls CheckSessionExistence
func Logout(c *gin.Context)  {
	var err error
	if _, err = CheckSessionExistence(c); err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Delete("123")
	if err = session.Save(); err != nil {
		c.HTML(
			http.StatusInternalServerError,
			"errors/error.html",
			gin.H{"error": err.Error()})
		return
	}

}
