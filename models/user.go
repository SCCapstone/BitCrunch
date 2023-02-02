// models.user.go

package models

import (
	"errors"
	"strings"
)

/*
user struct has a username and a password
*/
type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

/*
List of users
*/
var userList = []User{
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
func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := User{Username: username, Password: password}
	userList = append(userList, u)

	return &u, nil
}

/*
Check if the inputted username is available
*/
func isUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true
}
