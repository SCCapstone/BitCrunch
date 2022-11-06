package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashCost = 10 // min = 4, max = 31
)

type user struct {
	username string
	password []byte
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

	userr := &user{
		username: username,
		password: hash(password),
		email:    email,
	}

	// TODO
	// Add to db

	return userr, nil
}

func hash(password string) []byte {
	return bcrypt.GenerateFromPassword([]byte(password), hashCost)
}

func CheckValidPassword(username, password string) bool {
	return false // TODO
}

func checkUsername(username string) bool {
	return true // TODO later (check db)
}

func checkPassword(password string) bool {
	return true // TODO later
}

func checkEmail(email string) bool {
	return true // TODO later (probably will want to use regex)
}
