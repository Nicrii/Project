package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id int64 `json:"id"`
}

type PrivateUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (user *User) Marshall(isPublic bool) ([]byte, error) {
	if isPublic {
		return json.Marshal(PublicUser{Id: user.Id})
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return json.Marshal(privateUser)
}
