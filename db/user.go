package db

import (
	"context"
	"fmt"
	"log"
	"regexp"
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

func (u *user) IsAdmin() bool {
	return u.admin != 0
}

func (db *base) New(username, password, email string, administrator int) (*user, error) {
	if db.checkUsername(username) != nil {
		return nil, fmt.Errorf("Username \"%s\" is not unique. User creation failed.", username)
	}
	if checkPassword(password) != nil {
		return nil, fmt.Errorf("Password \"%s\" is not sufficient.", password)
	}
	if checkEmail(email) != nil {
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

/*
This function will check if a
password for a user is correct.
Returns true if the password
is valid for the user.
False otherwise.
*/
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

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return true
	}
	return false
}

/*
This function will check if
the username is available.
Returns nil if it is.
An error otherwise.
*/
func (db *dbase) checkUsername(username string) error {
	if !db.opened {
		db.Open() // will crash if this fails
	}
	query := "SELECT username FROM users"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	var test string
	row := stmt.QueryRowContext(ctx, username)
	if err = row.Scan(&test); err == nil {
		return fmt.Errorf("Username exists!")
	}
	return nil
}

/*
This function checks if
a password is allowed.
In other words, if the
password has the right length
and number of symbols, etc.
Returns nil if the password
is sufficient.
*/
func checkPassword(password string) error {
	// password must be at least 10 characters
	if len(password) < 10 {
		return fmt.Errorf("Password length=%d, need>=%d", len(password), 10)
	}
	// password must have at least one digit
	reg := regexp.MustCompile("\\d")
	if !reg.Match([]byte(password)) {
		return fmt.Errorf("Password doesn't have a digit.")
	}
	// password must have a symbol
	// symbol must be one of !@#$%^&*()
	reg = regexp.MustCompile("[!|@|#|$|$|%|^|7|*|(|)]")
	if !reg.Match([]byte(password)) {
		return fmt.Errorf("Password doesn't have a symbol.")
	}
	//password must have an uppercase character
	reg = regexp.MustCompile("[A-Z]")
	if !reg.Match([]byte(password)) {
		return fmt.Errorf("Password doesn't have an uppercase letter.")
	}
	//password must have a lowercase char
	reg = regexp.MustCompile("[a-z]")
	if !reg.Match([]byte(password)) {
		return fmt.Errorf("Password doesn't have a lowercase letter.")
	}
	return nil
}

/*
This function will check
to see that an email is valid
based on regex.
Returns nil if it is good to use.
*/
func checkEmail(email string) error {
	reg := regexp.MustCompile("(\\w+@[a-zA-Z_]+?\\.[a-zA-Z]{2,6})")
	if !reg.Match([]byte(email)) {
		return fmt.Errorf("Incorrect email!")
	}
	return nil
}

/*
Add a user to the user table in the DB.
Supply the username, password, and email for the user.
The 'admin' parameter may or may not be used
at this point. If used, admin != 0 means
the user IS an admin.
*/
func (db *dbase) AddUser(username, password, email string, admin int) error {
	if !db.opened {
		db.Open() // will crash if this fails
	}
	query := "INSERT INTO users(username, password, email, admin) VALUES (?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.sqldb.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, username, password, email, admin)
	if err != nil {
		return err
	}

	return nil
}
