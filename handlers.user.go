// handlers.user.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// these handle the specific pages startup (i.e. /login, /sign)
func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	print(username)
	print(password)

	if isUserValid(username, password) {
		render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func showSignUp(c *gin.Context) {
	render(c, gin.H{
		"title": "Sign Up"}, "signup.html")
}

func showDraggable(c *gin.Context) { // for testing draggable objects
	//handleDragPageStart(c) // testing this out
	render(c, gin.H{"title": "DragTest"}, "dragables.html")
}

func showSettings(c *gin.Context) {
	// cookies check would go here
	render(c, gin.H{
		"title": "Settings"}, "settings.html")
}

func showMap(c *gin.Context) {
	// cookies check would go here
	render(c, gin.H{
		"title": "Map"}, "map-base.html")
}
