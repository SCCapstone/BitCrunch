// main.go

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	middleware "github.com/SCCapstone/BitCrunch/middleware"
	// models "github.com/SCCapstone/BitCrunch/models"
	db "github.com/SCCapstone/BitCrunch/db"
	"github.com/gin-gonic/gin"
)

// Create the router
var router *gin.Engine

/*
Configures the router to load HTML templates
Sets the lower memory limit
Initializes the routes for the router
Hard codes the port for hosting
*/
func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.MaxMultipartMemory = 8 << 20
	InitializeRoutes()
	router.Run(":80")
}

/*
Properly renders template depending on format
*/
func Render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

/*
Initializes the routes for the entire project
*/
func InitializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(middleware.SetUserStatus())

	// Handle the index route
	router.GET("/", middleware.EnsureNotLoggedIn(), showLoginPage)

	// Group user related routes together
	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login, ensure user is not logged in using middleware
		// Render the login page
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), showLoginPage)

		// Handle POST requests at /u/login, ensure user is not logged in using middleware
		// Login the user
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), performLogin)

		// Handle GET requests at /u/logout, ensure user is logged in using middleware
		// Logout the user
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), logout)

		// Handle GET requests at /u/logout, ensure user is logged in using middleware
		// Display the logout modal
		userRoutes.GET("/logout_modal", middleware.EnsureLoggedIn(), display_logout_modal)

		// Handle GET requests at /u/add_layer_modal, ensure user is logged in using middleware
		// Display the add layer modal
		userRoutes.GET("/add_layer_modal", middleware.EnsureLoggedIn(), display_add_layer_modal)

		// Handle POST requests at /u/add_layer, ensure user is logged in using middleware
		// Add the layer
		// userRoutes.POST("/add_layer", middleware.EnsureLoggedIn(), AddLayer)

		// Handle POST requests at /u/view_layer, ensure user is logged in using middleware
		// Render the image to map
		userRoutes.POST("/view_layer", middleware.EnsureLoggedIn(), viewLayer)

		// Handle GET requests at /u/register, ensure user is not logged in using middleware
		//Render the registration page
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register, ensure user is not logged in using middleware
		//Register the user
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), register)
	}
	// Handle GET requests at /map, ensure user is logged in using middleware
	// Render the index page
	router.GET("/map", middleware.EnsureLoggedIn(), showMap)
}

/*
Renders the login page
*/
// NOTE: Moved from handlers.user.go
func showLoginPage(c *gin.Context) {
	Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

/*
Renders the registration page
*/
func showRegistrationPage(c *gin.Context) {
	Render(c, gin.H{
		"title": "Register"}, "register.html")
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

	if db.CheckValidPassword(username, password) == nil {
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		Render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")
	} else {
		fmt.Print("username:", db.CheckUsername(username))
		fmt.Print("password:", db.CheckPassword(password))
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
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

	if _, err := db.CreateUser(username, password,"temp@email.com", 1); err == nil {
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		Render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
	}
	}

/*
Clears the cookie and redirects to the home page
*/
func logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

/*
Renders the Logout Modal when the user presses the logout button
*/
func display_logout_modal(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"LogoutModal": "Logout Modal",
	})
}

/*
Renders the Add Layer Modal when the user presses the add layer button
*/
func display_add_layer_modal(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"AddLayerModal": "Add Layer Modal",
	})
}

/*
Generates a random 16 character string as the session token
*/
func GenerateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

/*
Gets the proper floor from the list of floors based on its name
Renders the proper floor image onto the map
*/
func viewLayer(c *gin.Context) {
	name := c.PostForm("layer")
	fmt.Println("here", name)
	// floors := models.GetAllFloors()
	// for i := 0; i < len(floors); i++ {
	// 	if floors[i].Name == name {
	// 		Render(c, gin.H{
	// 			"title":   "Map",
	// 			"payload": floors,
	// 			"Image":   "../" + floors[i].ImageFile,
	// 		}, "index.html")
	// 	}
	// }
}

/*
Renders the index with updated layer values
*/
func showMap(c *gin.Context) {

	fmt.Print(db.GetAllFloors())
	floors := db.GetAllFloors()

	for i := 0; i < len(floors); i++ {
		
	}
	


	Render(c, gin.H{
		"title":   "Map",
		"payload": floors,
	}, "index.html")
}

/*
Adds a layer with a layer name inputted from the user
Saves uploaded image to static/assets folder
Creates a new floor and adds it to the list of floors, calls showMap to render the map with updates
*/
// func AddLayer(c *gin.Context) {
// 	layer_name := c.PostForm("layer_name")

// 	file, err := c.FormFile("layer_image")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = c.SaveUploadedFile(file, "static/assets/"+file.Filename)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	models.CreateNewFloor(layer_name, "static/assets/"+file.Filename)
// 	showMap(c)
// }
