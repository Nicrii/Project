package services

import (
	"github.com/Nicrii/Project/users-api/domain/users"
	"github.com/Nicrii/Project/users-api/utils/crypto_utils"
	"github.com/Nicrii/Project/users-api/utils/errors"
)

var UserService userServiceInterface = &userService{}

type userService struct{}
type userServiceInterface interface {
	GetUser(userId int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Update(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}
