package domain

import (
	"fmt"
	"github.com/Project/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{Id: 1, FirstName: "FirstName", LastName: "LastName", Email: "sdfsdfsd@gmail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with id %v does not exist", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
