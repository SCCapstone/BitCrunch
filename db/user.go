package db

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	users    = "users.db"
	hashCost = 10 // min = 4, max = 31
)

type user struct {
	username string
	password []byte
	email    string
	admin    int
}

func CreateUser(username, password, email string, admin int) (user, error) {
	u := user{
		username: "",
		password: []byte(""),
		email:    "",
		admin:    0,
	}
	if CheckUsername(username) != nil {
		return u, fmt.Errorf("Username \"%s\" is already in use.", username)
	}
	if err := CheckPassword(password); err != nil {
		// Return specific error
		return u, err
	}
	if err := checkEmail(email); err != nil {
		// Return specific error
		return u, err
	}

	// Everything checks out so
	// return the user and
	// add to db
	u.username = username
	hashed, err := hash(password)
	if err != nil {
		return user{}, err
	}
	u.password = hashed
	u.email = email
	u.admin = admin

	if err := writeUser(u); err != nil {
		return user{}, fmt.Errorf("Failed to write user to db file!")
	}

	return u, nil
}

func writeUser(u user) error {
	fil, err := os.OpenFile(users, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fil.Close()
	// Creating the string from the user details
	// to append to the file
	writeString := fmt.Sprintf("%s\t%s\t%s\t%d\n", u.username, u.password, u.email, u.admin)
	_, err = fil.WriteString(writeString)
	if err != nil {
		return err
	}
	// Finally, no problems
	return nil
}

/*
Attempts to find the user
in the db file and then
return a user struct from the
read data. Returns error
if the user is not found.
*/
func ReadUser(uname string) (u user, err error) {
	fi, err := os.Open(users)
	if err != nil {
		return
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	// Reading line-by-line to find the username
	for scan.Scan() {
		line = strings.Split(scan.Text(), "\t")
		if line[0] == uname {
			// User found, creating it to return
			u = user{
				username: line[0],
				password: []byte(line[1]),
				email:    line[2],
				admin:    0,
			}
			return u, nil
		}
	}
	// User was not found in the file
	return user{}, fmt.Errorf("User not found.")
}

/*
Deletes a user from the db
file. Returns nil if successful.
An error otherwise.
*/
func DeleteUser(uname string) error {
	// Creating a temp file
	delMe, err := os.Create(fmt.Sprintf("temp%s.tmp", uname))
	if err != nil {
		return err
	}
	fi, err := os.Open(users)
	if err != nil {
		return err
	}
	scan := bufio.NewScanner(fi)
	var line string
	for scan.Scan() {
		line = scan.Text()
		if strings.Split(line, "\t")[0] != uname {
			delMe.WriteString(line + "\n")
		}
	}
	// Done with the main file
	// Removing it
	fi.Close()
	err = os.Remove(users)
	if err != nil {
		return err
	}

	// Renaming the file without the
	// floor to be deleted to the users.db
	delMe.Close()
	RenameFile(delMe.Name())

	// Done, clean up
	return nil
}

func RenameFile(filename string) error {
	err := os.Rename(filename, users)
	if err != nil {
		return err
	}
	return nil
}

/*
This function will check if a
password for a user is correct.
Returns nil if the password
is valid for the user.
Error otherwise.
*/
func CheckValidPassword(username, passwd string) (err error) {
	userr, err := ReadUser(username)
	if err != nil {
		return
	}
	pass := []byte(passwd)
	// Comparing the db hash and the supplied hash
	return bcrypt.CompareHashAndPassword(userr.password, pass)
}

/*
This function will check if
the username is available.
Returns nil if it is.
An error otherwise.
*/
func CheckUsername(u string) error {
	fi, err := open(users)
	if err != nil {
		return err
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	for scan.Scan() {
		line = strings.Split(scan.Text(), "\t")
		if line[0] == u {
			return fmt.Errorf("Username found!")
		}
	}
	// Username not found
	return nil
}

/*
Local function used to open
a specific file and return it.
Ensures that the file is not
already created.
*/
func open(file string) (fi *os.File, err error) {
	fi, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return
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
func CheckPassword(password string) error {
	// password must be at least 10 characters
	if len(password) < 10 {
		return fmt.Errorf("Password needs to have at least 10 characters.")
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
func checkEmail(e string) error {
	// Checking for valid email
	reg := regexp.MustCompile("(\\w+@[a-zA-Z_]+?\\.[a-zA-Z]{2,6})")
	if !reg.Match([]byte(e)) {
		return fmt.Errorf("Please enter a valid email address.")
	}
	// Checking if email address already in use
	fi, err := open(users)
	if err != nil {
		return err
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	for scan.Scan() {
		line = strings.Split(scan.Text(), "\t")
		if line[2] == e {
			return fmt.Errorf("Email is already in use with another account.")
		}
	}
	// Email ok
	return nil
}

func hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), hashCost)
}
