package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type UserSession struct {
	ID         uint64    `gorm:"primary_key"`
	SessionKey string    `gorm:"size:255;unique;not null"`
	ExpireAt   time.Time `gorm:"size:255;unique;not null"`
	UserId     uint64
}

// FindSessionByKey allows you to find a specific session using its sessionKey.
// Also preload the associated user.
func FindSessionByKey(SessionKey string) (*UserSession, error) {
	var err error
	var userSession UserSession
	err = db.Debug().Where("session_key = ?", SessionKey).Find(&userSession).Preload("Users").Error
	if err != nil {
		return &UserSession{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserSession{}, errors.New("Session Not Found")
	}
	fmt.Println(&userSession.UserId)
	return &userSession, nil
}

// CreateUserSession creates an user_session row in database
func CreateUserSession(user *User, sessionKey string) (*UserSession, error) {
	var err error
	var userSession UserSession
	userSession.SessionKey = sessionKey
	userSession.ExpireAt = time.Now().Add(20 * time.Minute)
	userSession.UserId = user.ID
	err = db.Debug().Create(&userSession).Error

	if err != nil {
		return nil, err
	}
	return &userSession, nil
}

// EditUserSessionByKey update an user_session row in database from it session key
func EditUserSessionByKey(session *UserSession, sessionKey string) (*UserSession, error) {
	var err error
	var oldSession UserSession
	err = db.Debug().Where("session_key = ?", sessionKey).First(&oldSession).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("Session Not Found")
	}
	session.SessionKey = uuid.NewV4().String()
	session.ExpireAt = time.Now().Add(20 * time.Minute)
	err = db.Debug().Save(&session).Error
	if err != nil {
		return nil, errors.New("Could'nt update session")
	}

	return session, nil
}

func FindSessionByUserId(id uint64) (*UserSession, error) {
	var err error
	var userSession UserSession
	err = db.Debug().Where("user_id = ?", id).Find(&userSession).Error
	if err != nil {
		return &UserSession{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserSession{}, errors.New("Session Not Found")
	}
	return &userSession, nil
}
