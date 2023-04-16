// application.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	middleware "github.com/SCCapstone/BitCrunch/middleware"
	// models "github.com/SCCapstone/BitCrunch/models"
	db "github.com/SCCapstone/BitCrunch/db"
	rd "github.com/SCCapstone/BitCrunch/scriptrunner"
	"github.com/gin-gonic/gin"
)

// Create the router
var router *gin.Engine

var currentFloor = ""
var currentFile = ""
var currentDevice = ""

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
	router.Run(":5000")
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

		userRoutes.POST("/logout", middleware.EnsureLoggedIn(), logout)

		// Handle GET requests at /u/logout, ensure user is logged in using middleware
		// Display the logout modal
		userRoutes.GET("/logout_modal", middleware.EnsureLoggedIn(), displayModal("LogoutModal", "Logout Modal"))

		// Handle GET requests at /u/add_layer_modal, ensure user is logged in using middleware
		// Display the add layer modal
		userRoutes.GET("/add_layer_modal", middleware.EnsureLoggedIn(), displayModal("AddLayerModal", "Add Layer Modal"))

		// Handle POST requests at /u/add_layer, ensure user is logged in using middleware
		// Add the layer
		userRoutes.POST("/add_layer", middleware.EnsureLoggedIn(), AddLayer)

		// Handle GET requests at /u/add_device_modal, ensure user is logged in using middleware
		// Display the add device modal
		userRoutes.GET("/add_device_modal", middleware.EnsureLoggedIn(), displayModal("AddDeviceModal", "Add Device Modal"))

		// Handle POST requests at /u/add_device, ensure user is logged in using middleware
		// Add the device
		userRoutes.POST("/add_device", middleware.EnsureLoggedIn(), AddDevice)

		// Handle GET requests at /u/delete_layer_modal, ensure user is logged in using middleware
		// Display the delete layer modal
		userRoutes.GET("/delete_layer_modal", middleware.EnsureLoggedIn(), displayModal("DeleteLayerModal", "Delete Layer Modal"))

		// Handle POST requests at /u/delete_layer, ensure user is logged in using middleware
		// Delete the layer
		userRoutes.POST("/delete_layer", middleware.EnsureLoggedIn(), DeleteLayer)

		// Handle GET requests at /u/edit_layer_modal, ensure user is logged in using middleware
		// Display the edit layer modal
		userRoutes.GET("/edit_layer_modal", middleware.EnsureLoggedIn(), displayModal("EditLayerModal", "Edit Layer Modal"))

		userRoutes.POST("/edit_layer", middleware.EnsureLoggedIn(), EditLayer)
		// Handle POST requests at /u/view_layer, ensure user is logged in using middleware
		// Render the image to map
		userRoutes.POST("/view_layer", middleware.EnsureLoggedIn(), viewLayer)

		userRoutes.POST("/view_device", middleware.EnsureLoggedIn(), viewDevice)

		// Handle GET requests at /u/register, ensure user is not logged in using middleware
		//Render the registration page
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), showRegistrationPage)

		// Handle POST requests at /u/register, ensure user is not logged in using middleware
		//Register the user
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), register)

		userRoutes.GET("/delete_account_modal", middleware.EnsureLoggedIn(), displayModal("DeleteAccountModal", "Delete Account Modal"))

		userRoutes.GET("/delete_account", middleware.EnsureLoggedIn(), delete_account)

		userRoutes.GET("/delete_device_modal", middleware.EnsureLoggedIn(), displayModal("DeleteDeviceModal", "Delete Device Modal"))

		userRoutes.GET("/delete_device", middleware.EnsureLoggedIn(), deleteDevice)

		userRoutes.GET("/run_script", middleware.EnsureLoggedIn(), displayModal("ScriptModal", "Script Modal"))

		userRoutes.GET("/ping_device", middleware.EnsureLoggedIn(), pingDevice)

		userRoutes.POST("/edit_device", middleware.EnsureLoggedIn(), editDevice)
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
		c.SetCookie("current_user", username, 3600, "/", "localhost", false, true)

		Render(c, gin.H{
			"title": "Successful Login"}, "login-successful.html")
	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

/*
Obtains user inputted username, password,
Password confirmation, and email
If the user is properly created, set the token in a cookie
Log the user in by rendering successful login
If the user created is invalid, renders an error
*/
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")
	email := c.PostForm("email")

	if password != confirm_password {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": fmt.Sprintf("Passwords \"%s\" and \"%s\" do not match.", password, confirm_password)})
	} else if _, err := db.CreateUser(username, password, email, 1); err == nil {
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)
		c.SetCookie("current_user", username, 3600, "/", "localhost", false, true)

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

func displayModal(modalName string, msg string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			modalName: msg,
		})
	}
	return gin.HandlerFunc(fn)
}

func renderError(c *gin.Context, modal, modalName, errorTitle, et, errorMessage, emsg string) {
	c.HTML(http.StatusBadRequest, "index.html", gin.H{
		modal:        modalName,
		errorTitle:   et,
		errorMessage: emsg})
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
	imageName := ""
	floors, _ := db.GetAllFloors()
	floorNames := []string{}
	deviceNames := []string{}
	scriptNames := []string{}
	for i := 0; i < len(floors); i++ {
		str := fmt.Sprintf("%#v", floors[i])
		comma := strings.Index(str, ",")
		substr := str[15 : comma-1]
		floorNames = append(floorNames, substr)
	}
	for i := 0; i < len(floorNames); i++ {
		if floorNames[i] == name {
			fileIO, err := os.OpenFile("devices/"+name+".txt", os.O_RDWR, 0600)
			if err != nil {
				fmt.Println(err)
			}
			defer fileIO.Close()
			rawBytes, err := ioutil.ReadAll(fileIO)
			if err != nil {
				fmt.Println(err)
			}
			lines := strings.Split(string(rawBytes), "\n")
			for i, line := range lines {
				if i == 0 {
					imageName = line
				}
			}
		}
	}

	setCurrentFloor(name)
	setCurrentFile(imageName)

	devices, _ := db.GetAllDevicesForFloor(getCurrentFloor())

	for i := 0; i < len(devices); i++ {
		str := fmt.Sprintf("%#v", devices[i])
		comma := strings.Index(str, ",")
		substr := str[16 : comma-1]
		deviceNames = append(deviceNames, substr)
	}

	files, err := ioutil.ReadDir("/static/assets")
	if err != nil {
		fmt.Println(err)
	}

	for _, files := range files {
		temp := files.Name()
		if strings.Contains(temp, "script.txt") {
			delete := strings.Index(temp, "script.txt")
			scriptName := temp[0 : delete-1]
			scriptNames = append(scriptNames, scriptName)
		}
	}

	Render(c, gin.H{
		"title":           "Map",
		"payload":         floorNames,
		"Image":           "static/assets/" + imageName,
		"EditLayerButton": "EditLayerButton",
		"devices":         deviceNames,
		"scripts":         scriptNames,
	}, "index.html")
}

func viewDevice(c *gin.Context) {
	name := c.PostForm("device")
	setCurrentDevice(name)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"ViewDeviceModal": "ViewDeviceModal",
		"DeviceName":      name,
		"DeviceIP":        db.GetIP(name),
	})
}

// Runs scripts
func RunScript(c *gin.Context) {
	script := c.PostForm("script")
	IP := c.PostForm("ip")
	script_file := script + "script.txt"

	if err, script := rd.RunFromScript(script_file, IP); err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"ScriptOutputModalError": "Script Output Modal Error",
			"ErrorTitle":             "Script Error",
			"ErrorMessage":           err.Error()})
		return
	} else {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"ScriptOutputModal": "Script Output Modal",
			"ScriptOutput":      script})
		return
	}
}

/*
Renders the index with updated layer values
*/
func showMap(c *gin.Context) {
	floors, _ := db.GetAllFloors()
	floorNames := []string{}

	for i := 0; i < len(floors); i++ {
		str := fmt.Sprintf("%#v", floors[i])
		comma := strings.Index(str, ",")
		substr := str[15 : comma-1]
		floorNames = append(floorNames, substr)
	}

	Render(c, gin.H{
		"title":   "Map",
		"payload": floorNames,
	}, "index.html")
}

/*
Adds a layer with a layer name inputted from the user
Saves uploaded image to static/assets folder
Creates a new floor and adds it to the list of floors, calls showMap to render the map with updates
*/
func AddLayer(c *gin.Context) {
	layer_name := c.PostForm("layer_name")
	file, err := c.FormFile("layer_image")
	if err != nil {
		renderError(c, "AddLayerModal", "Add Layer Modal", "ErrorTitle", "Add Layer Failed", "ErrorMessage", "Image file could not be found.")
		return
	}
	err = c.SaveUploadedFile(file, "static/assets/"+file.Filename)
	if err != nil {
		renderError(c, "AddLayerModal", "Add Layer Modal", "ErrorTitle", "Failed to Add Layer", "ErrorMessage", "Image file could not be saved.")
		return
	}

	if _, err := db.CreateFloor(layer_name, layer_name+".txt"); err != nil {
		renderError(c, "AddLayerModal", "Add Layer Modal", "ErrorTitle", "Failed to Add Layer", "ErrorMessage", err.Error())
		return
	} else {
		createDeviceFile(layer_name, file.Filename)
		showMap(c)
	}
}

/*
Edit the name, image, or both of the current layer
*/
func EditLayer(c *gin.Context) {
	old_layer_name := getCurrentFloor()
	// old_file_name := getCurrentFile()
	layer_name := c.PostForm("layer_name")
	// fname := old_file_name
	if len(layer_name) == 0 {
		layer_name = old_layer_name
	}
	file, err := c.FormFile("layer_image")
	if err != nil {
		renderError(c, "EditLayerModal", "Edit Layer Modal", "ErrorTitle", "Failed to Edit Layer", "ErrorMessage", "Image file could not be found.")
		return
	} else {
		err = c.SaveUploadedFile(file, "static/assets/"+file.Filename)
		// fname = file.Filename
		if err != nil {
			renderError(c, "EditLayerModal", "Edit Layer Modal", "ErrorTitle", "Failed to Edit Layer", "ErrorMessage", "Image file could not be saved.")
			return
		}
	}

	if err := db.DeleteFloor(old_layer_name); err != nil {
		renderError(c, "EditLayerModal", "Edit Layer Modal", "ErrorTitle", "Failed to Edit Layer", "ErrorMessage", err.Error())
		return
	}

	if _, err := db.CreateFloor(layer_name, layer_name+".txt"); err != nil {
		renderError(c, "EditLayerModal", "Edit Layer Modal", "ErrorTitle", "Failed to Edit Layer", "ErrorMessage", err.Error())
		return
	}

	renameDeviceFile(old_layer_name, layer_name)
	saveNewImage(file.Filename, layer_name)

	showMap(c)
}

/*
Adds a device with a device name inputted from the user
Saves uploaded image to static/assets folder
adds the device to the floor's deviceList file
*/
func AddDevice(c *gin.Context) {
	device_name := c.PostForm("device_name")
	device_ip := c.PostForm("device_ip")
	device_image, err := c.FormFile("device_image")

	if err != nil {
		renderError(c, "AddDeviceModal", "Add Device Modal", "ErrorTitle", "Failed to Add Device", "ErrorMessage", "Image file could not be found.")
		return
	}
	err = c.SaveUploadedFile(device_image, "static/assets/"+device_image.Filename)
	if err != nil {
		renderError(c, "AddDeviceModal", "Add Device Modal", "ErrorTitle", "Failed to Add Device", "ErrorMessage", "Image file could not be saved.")
		return
	}

	if _, err := db.CreateDevice(device_name, device_ip, "static/assets/"+device_image.Filename, getCurrentFloor()); err != nil {
		renderError(c, "AddDeviceModal", "Add Device Modal", "ErrorTitle", "Failed to Add Device", "ErrorMessage", err.Error())
		return
	}
	showMap(c)
}

func deleteDevice(c *gin.Context) {
	name := getCurrentDevice()
	floor := getCurrentFloor()
	db.DeleteDevice(name, floor)
	showMap(c)
}

func editDevice(c *gin.Context) {
	floor := getCurrentFloor()
	name := getCurrentDevice()
	newName := c.PostForm("device_name")
	newIP := c.PostForm("device_ip")
	newImage, err := c.FormFile("device_image")
	if newImage != nil {
		err = c.SaveUploadedFile(newImage, "static/assets/"+newImage.Filename)
	}
	// checking IP is valid
	if (len(newIP) > 0) && (newIP != db.GetIP(name)) {
		if err := db.CheckIP(newIP); err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"ViewDeviceModal": "ViewDeviceModal",
				"DeviceName":      name,
				"DeviceIP":        db.GetIP(name),
				"ErrorTitle":      "Failed to Edit Device",
				"ErrorMessage":    err.Error(),
			})
			return
		} else {
			db.EditDevice(name, name, newIP, db.GetImage(name), floor)
		}
	}
	// adding image if present
	if newImage != nil {
		db.EditDevice(name, name, db.GetIP(name), "static/assets/"+newImage.Filename, floor)
	}
	// checking device name is unique for floor
	if (len(newName) > 0) && (newName != name) {
		if err = db.CheckDevice(newName, floor); err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"ViewDeviceModal": "ViewDeviceModal",
				"DeviceName":      name,
				"DeviceIP":        db.GetIP(name),
				"ErrorTitle":      "Failed to Edit Device",
				"ErrorMessage":    err.Error(),
			})
			return
		} else {
			db.EditDevice(name, newName, db.GetIP(name), db.GetImage(name), floor)
		}
	}
	showMap(c)
}

/*
Deletes a layer from the list of floors,
calls showMap to render the map with updates
*/
func DeleteLayer(c *gin.Context) {
	name := getCurrentFloor()
	if err := db.DeleteFloor(name); err != nil {
		renderError(c, "DeleteLayerModal", "Delete Device Modal", "ErrorTitle", "Failed to Add Device", "ErrorMessage", err.Error())
		return
	}
	removeDeviceFile("devices/" + name + ".txt")
	showMap(c)
}

func pingDevice(c *gin.Context) {
	device := getCurrentDevice()
	ip := db.GetIP(device)

	var waitgroup sync.WaitGroup
	waitgroup.Add(1) // used for waiting on all waits to finish (but we only use one)
	var output string
	go func() {
		_, output = rd.RunFromScript("pingscript.txt", ip) // actually run the script
		waitgroup.Done()                                   // once GetIP is finished, reduce the waitgroup counter by one
	}()
	waitgroup.Wait() // blocks until all wait groups are at 0 (so once)

	// displayModal("ScriptModal", "Script Modal")

	// HERE: This call is the issue in the AWS website.
	// ask for explanation on how this works during improptu meeting 4/16/23 (see edit device)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Output": output,
	})

}

func createDeviceFile(name string, filename string) {
	file, err := os.OpenFile("devices/"+name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writeString := fmt.Sprintf(filename + "\n")
	_, err = file.WriteString(writeString)
}

func removeDeviceFile(name string) {
	err := os.Remove(name)
	if err != nil {
		log.Fatal(err)
	}
}

func renameDeviceFile(old, new string) {
	os.Rename("devices/"+old+".txt", "devices/"+new+".txt")
}

func saveNewImage(new_image, layer string) {
	fi, err := ioutil.ReadFile("devices/" + layer + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(fi), "\n")

	for i := range lines {
		if i == 0 {
			lines[i] = new_image
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("devices/"+layer+".txt", []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func delete_account(c *gin.Context) {
	logout(c)
	current_user, _ := c.Cookie("current_user")
	db.DeleteUser(current_user)
}

func setCurrentFloor(floorName string) {
	if len(floorName) > 0 {
		currentFloor = floorName
	}
}

func getCurrentFloor() (floorName string) {
	return currentFloor
}

func setCurrentFile(fileName string) {
	if len(fileName) > 0 {
		currentFile = fileName
	}
}

func getCurrentFile() (fileName string) {
	return currentFile
}

func setCurrentDevice(deviceName string) {
	if len(deviceName) > 0 {
		currentDevice = deviceName
	}
}

func getCurrentDevice() (deviceName string) {
	return currentDevice
}
