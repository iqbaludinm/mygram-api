package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/iqbaludinm/mygram-api/models"
	"gorm.io/gorm"
)

// interface
type UserRepository interface {
	RegisterUser(user models.User) (res models.User, err error)
	LoginUser(user models.LoginUser) (res models.User, err error)
}

func (r BaseRepository) RegisterUser(user models.User) (res models.User, err error) {
	var existingUser models.User
	err = r.gorm.Where("username = ? OR email = ?", user.Username, user.Email).Find(&existingUser).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}
	if existingUser.ID != 0 {
		if existingUser.Username == user.Username {
			return res, fmt.Errorf("User with username '%s' already exists", user.Username)
		} else if existingUser.Email == user.Email {
			return res, fmt.Errorf("User with email '%s' already exists", user.Email)
		}
	}

	err = r.gorm.Create(&user).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r BaseRepository) LoginUser(user models.LoginUser) (res models.User, err error) {
	err = r.gorm.Debug().Where("email = ?", user.Email).Take(&res).Error
	if err != nil {
		log.Println(err) 
		return res, errors.New("Invalid Email or Password!")
	}
	return res, nil
}