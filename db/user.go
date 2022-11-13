package db

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashCost = 10 // min = 4, max = 31
)

type user struct {
	username string
	password []byte
	email    string
	admin    int
}

func (u *user) IsAdmin() int {
	return u.admin
}

func New(username, password, email string, administrator int) (*user, error) {
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
		admin:    administrator,
	}

	return userr, nil
}

func hash(password string) []byte {
	return bcrypt.GenerateFromPassword([]byte(password), hashCost)
}

func (db *dbase) CheckValidPassword(username, password string) bool {
	if !db.opened {
		db.Open()
	}
	query := "SELECT password FROM users WHERE username = ?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.sqldb.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	var hash string
	row := stmt.QueryRowContext(ctx, username)
	if err := row.Scan(&hash); err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return false
	}
	return true
}

func checkUsername(username string) bool {
	return true // TODO
}

func checkPassword(password string) bool {
	return true // TODO later
}

func checkEmail(email string) bool {
	return true // TODO later (probably will want to use regex)
}
