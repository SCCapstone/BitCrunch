// routes.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"login.html",
			gin.H{
				"title": "Login Page",
			},
		)

	})

	router.POST("/login", performLogin)

	router.GET("/sign", showSignUp)
}
