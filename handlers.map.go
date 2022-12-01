// handlers.map.go

package main

import (
	"github.com/gin-gonic/gin"
)

func showMap(c *gin.Context) {
	floors := getAllFloors()
	

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Map",
		"payload": floors}, "index.html")
}

func addLayer(c *gin.Context) {
	layer_name := c.PostForm("layer_name")
	layer_image := c.PostForm("layer_image")
	createNewFloor(layer_name, "mock.txt")
	showMap(c)
}