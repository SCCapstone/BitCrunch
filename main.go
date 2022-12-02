// main.go

package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//Create the router
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
	initializeRoutes()
	router.Run(":80")
}

/*
Properly renders template depending on format
*/
func render(c *gin.Context, data gin.H, templateName string) {
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
