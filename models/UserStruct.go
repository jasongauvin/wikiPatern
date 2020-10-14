package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint64      `gorm:"primary_key"`
	Uuid 	 string
	Name     string      `gorm:"size:255"`
	Password string      `gorm:"size:255"`
	Email    string      `gorm:"size:255"`
	Admin    bool        `gorm:"default:false"`
}

// Hash the password passed as argument
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Verify the hashed password and the password to check password consistency
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateUuid(user *User)  {
	uid := uuid.NewV4().String()
	user.Uuid = uid
}
// Get the user model and hash it password before saving it
func BeforeSave(user *User) error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	GenerateUuid(user)
	return nil
}

// CreateUser creates an user row in database
func CreateUser(user *User) (*User, error) {
	var err error
	err = BeforeSave(user)
	if err != nil {
		return &User{}, err
	}
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

// EditUserByID update a user from its Id.
func EditUserByID(user *User) (*User, error) {
	var err error
	err = db.Debug().Save(&user).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return user, err
}

// DeleteUserByID delete a user from its Id.
func DeleteUserByID(id uint64) (User, error) {
	var err error
	var user User

	err = db.Debug().Delete(&user, id).Error
	if err != nil {
		return User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return User{}, errors.New("User Not Found")
	}
	return user, err
}

// FindUsers returns you a list of all stored users.
func FindUsers() (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// FindUsersById returns you a user from its Id.
func FindUserByID(uid uint64) (*User, error) {
	var err error
	var user User
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &user, err
}

// FindUsersById returns you a user from its Id.
func FindUserByUuiD(uuid string) (*User, error) {
	var err error
	var user User
	err = db.Debug().Model(User{}).Where("uuid = ?", uuid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &user, err
}

// FindUsersByEmail returns you a user from its Email.
func FindUserByEmail(email string) (*User, error) {
	var err error
	var user User
	err = db.Debug().Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return &user, err
}
