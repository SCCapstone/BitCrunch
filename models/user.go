// models.user.go

package models

import (
	"errors"
	"strings"

	db "github.com/SCCapstone/BitCrunch/db"
)

/*
user struct has a username and a password
*/
type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

/*
List of users
*/
var userList = []user{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
	{Username: "user3", Password: "pass3"},
}

/*
Checks if the username and password combination is valid
*/
func IsUserValid(username, password string) bool {
	for _, u := range userList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

/*
Registers a new user with given username/password by adding to list
*/
func RegisterNewUser(username, password, confirm_password, email string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if db.CheckUsername(username) != nil {
		return nil, errors.New("The username isn't available")
	} else if db.CheckPassword(password) != nil {
		return nil, db.CheckPassword(password)
	} else if password != confirm_password {
		return nil, errors.New("Passwords do not match")
	} else if db.checkEmail(email) != nil {
		return nil, errors.New("Email is not valid")
	}

	u := user{Username: username, Password: password}
	userList = append(userList, u)

	return &u, nil
}
