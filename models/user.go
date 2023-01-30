// models.user.go

package models

import (
	"errors"
	"regexp"
	"strings"
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
func RegisterNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
	} else if !isPasswordValid(password) {
		return nil, errors.New("Password must have at least 10 characters, consisting of a capital letter, lowercase letter, symbol, and number")
	}

	u := user{Username: username, Password: password}
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

/*
Check if the inputted password is valid
*/
func isPasswordValid(password string) bool {
	validChar, err := regexp.MatchString("[!@#$%^&* | A-z0-9]", password)
	if err == nil && len(password) >= 10 && validChar {
		return true
	}
	return false
}
