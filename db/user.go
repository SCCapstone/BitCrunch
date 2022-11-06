package db

import "fmt"

type user struct {
	username string
	password string
	email    string
}

func New(username, password, email string) (*user, error) {
	if !checkUsername(username) {
		return nil, fmt.Errorf("Username \"%s\" is not unique. User creation failed.", username)
	}
	if !checkPassword(password) {
		return nil, fmt.Errorf("Password \"%s\" is not sufficient.", password)
	}
	if !checkEmail(email) {
		return nil, fmt.Errorf("Email \"%s\" is either already used or not valid.", email)
	}

	// TODO
	// Compute hash of password and make that the password!!!!

	userr := &user{
		username: username,
		password: password,
		email:    email,
	}

	// TODO
	// Add to db

	return userr, nil
}

func checkUsername(username string) bool {
	return true // TODO later (actually check)
}

func checkPassword(password string) bool {
	return true // TODO later
}

func checkEmail(email string) bool {
	return true // TODO later (probably will want to use regex)
}
