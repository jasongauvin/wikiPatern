package services

import (
	"fmt"
	"github.com/jasongauvin/wikiPattern/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type RegisterForm struct {
	Email 	 string `form:"userEmail" binding:"required"`
	Password string `form:"userPassword" binding:"required"`
}

type LoginForm struct {
	Email 	 string `form:"userEmail" binding:"required"`
	Password string `form:"userPassword" binding:"required"`
}

// SaveUser creates a user.
// Return the created object
func SaveUser(email string, password string) (*models.User, error)  {
	var err error
	var user models.User
	user.Email = email
	user.Password = password

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
func AuthenticateUser(email string, password string) (*models.UserSession, error) {
	var err error
	var user *models.User

	user, err = models.FindUserByEmail(email)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("error:", err)
		return nil, err
	}
	userSession, err := CheckSessionExpiration(user.ID)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return userSession, nil
}

// CheckSessionExistence let you find the token by the sessionKey
// Return the UserSession
func CheckSessionExistence(sessionKey string) (*models.UserSession, error) {
	var userSession *models.UserSession
	var err error
	userSession, err = models.FindSessionByKey(sessionKey)
	if err != nil {
		return nil, err
	}

	return userSession, nil
}

// CheckSessionExpiration let you find the token by the user id
// Check if the UserSession is expired
// Update the UserSession with a new expire date and sessionToken
func CheckSessionExpiration(id uint64) (*models.UserSession, error) {
	userSession, err := models.FindSessionByUserId(id)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	if userSession.ExpireAt.After(time.Now()){
		var updatedSession *models.UserSession
		updatedSession, err = models.EditUserSessionByKey(userSession, userSession.SessionKey)
		return updatedSession, nil
	}
	return userSession, nil

}