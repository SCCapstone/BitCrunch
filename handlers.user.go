// handlers.user.go

package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

/*
Renders the login page
*/
func showLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

/*
Obtains user inputted username and password
Checks if the username/password combination is valid
If valid, setss token in a cookie
Renders successful login
If invalid, renders an error
*/
func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if isUserValid(username, password) {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

/*
Generates a random 16 character string as the session token
*/
func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

/*
Renders the Logout Modal when the user presses the logout button
*/
func display_logout_modal(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"LogoutModal":   "Logout Modal",
		})
}

/*
Renders the Add Layer Modal when the user presses the add layer button
*/
func display_add_layer_modal(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"AddLayerModal":   "Add Layer Modal",
		})
}

/*
Clears the cookie and redirects to the home page
*/
func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

/*
Renders the registration page
*/
func showRegistrationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Register"}, "register.html")
}

/*
Obtains user inputted username and password
If the user is properly created, set the token in a cookie
Log the user in by rendering successful login
If the user created is invalid, renders an error
*/
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if _, err := registerNewUser(username, password); err == nil {
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
}