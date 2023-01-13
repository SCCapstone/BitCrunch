// models.user.go

package main

import (
	"errors"
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
	user{Username: "user1", Password: "pass1"},
	user{Username: "user2", Password: "pass2"},
	user{Username: "user3", Password: "pass3"},
}

<<<<<<< HEAD
// logic check for user validation
=======
/*
Checks if the username and password combination is valid
*/
>>>>>>> main
func isUserValid(username, password string) bool {
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
func registerNewUser(username, password string) (*user, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	} else if !isUsernameAvailable(username) {
		return nil, errors.New("The username isn't available")
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