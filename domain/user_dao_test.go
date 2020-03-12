package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)
	assert.Nil(t, user, "Unexpected user with id 0")
	assert.NotNil(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user with id 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "FirstName", user.FirstName)
	assert.EqualValues(t, "LastName", user.LastName)
	assert.EqualValues(t, "sdfsdfsd@gmail.com", user.Email)
}
