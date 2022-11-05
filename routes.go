// routes.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// "plots" out all routes that can be taken, sets up the functions w them
func initializeRoutes() {
	router.GET("/", func(c *gin.Context) { // the / means on initial connection (welcome page)
		c.HTML(
			http.StatusOK,
			"login.html",
			gin.H{
				"title": "Login Page",
			},
		)

	}) // basic setup

	router.POST("/login", performLogin) // throw up performLogin AT (---)/login

	router.GET("/sign", showSignUp) // throw up showSign AT (---)/sign
}
