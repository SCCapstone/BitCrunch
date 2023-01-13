// routes.go

package main

<<<<<<< HEAD
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

	router.GET("/dragtest", showDraggable) //

	router.GET("/settings", showSettings) // for settings formatting

	router.GET("/map", showMap) // for map formatting

}
=======
/*
Initializes the routes for the entire project
*/
func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	router.GET("/", ensureNotLoggedIn(), showLoginPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login, ensure user is not logged in using middleware
		// Render the login page
		userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /u/login, ensure user is not logged in using middleware
		// Login the user
		userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout, ensure user is logged in using middleware
		// Logout the user
		userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// Handle GET requests at /u/logout, ensure user is logged in using middleware
		// Display the logout modal
		userRoutes.GET("/logout_modal", ensureLoggedIn(), display_logout_modal)

		// Handle GET requests at /u/add_layer_modal, ensure user is logged in using middleware
		// Display the add layer modal
		userRoutes.GET("/add_layer_modal", ensureLoggedIn(), display_add_layer_modal)

		// Handle POST requests at /u/add_layer, ensure user is logged in using middleware
		// Add the layer
		userRoutes.POST("/add_layer", ensureLoggedIn(), addLayer)

		// Handle POST requests at /u/view_layer, ensure user is logged in using middleware
		// Render the image to map
		userRoutes.POST("/view_layer", ensureLoggedIn(), viewLayer)

		// Handle GET requests at /u/register, ensure user is not logged in using middleware
		//Render the registration page
		userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register, ensure user is not logged in using middleware
		//Register the user
		userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}
	// Handle GET requests at /map, ensure user is logged in using middleware
	// Render the index page
	router.GET("/map", ensureLoggedIn(), showMap)
}
>>>>>>> main
