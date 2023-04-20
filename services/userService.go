package services

import (
	"errors"

	"github.com/iqbaludinm/mygram-api/helpers"
	"github.com/iqbaludinm/mygram-api/models"
)

// interface
type UserService interface {
	RegisterUser(user models.RegisterUser) (res models.User, err error)
	LoginUser(user models.LoginUser) (res string, err error)
}

func (s *BaseService) RegisterUser(user models.RegisterUser) (res models.User, err error) {
	users := models.User{
		Username: user.Username,
		Email: user.Email,
		Password: user.Password,
		Age: user.Age,
	}
	return s.repo.RegisterUser(users)
}

func (s *BaseService) LoginUser(user models.LoginUser) (res string, err error) {
	result, err := s.repo.LoginUser(user)
	if err != nil {
		return res, err
	}
	isValid := helpers.ComparePass([]byte(result.Password), []byte(user.Password))
	if !isValid {
		return res, errors.New("Invalid Email or Password!")
	}

	token := helpers.GenerateToken(result.ID, result.Email)
	
	return token, nil
}
