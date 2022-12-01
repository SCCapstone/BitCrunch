// handlers.map.go

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
)

func showMap(c *gin.Context) {
	floors := getAllFloors()
	

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Map",
		"payload": floors,
		}, "index.html")
}

func addLayer(c *gin.Context) {
	layer_name := c.PostForm("layer_name")

	file, err := c.FormFile("layer_image")
	if err != nil {
		log.Println(err)
	}
	log.Println(file.Filename)

	err = c.SaveUploadedFile(file, "static/assets/" + file.Filename)
	if err != nil {
		log.Println(err)
	}

	createNewFloor(layer_name, "static/assets/" + file.Filename)
	fmt.Println(getAllFloors())
	showMap(c)
}

func viewLayer(c *gin.Context) {
	name := c.PostForm("l_name")
	println("name: " + name)
	floors := getAllFloors()
	for i := 0; i < len(floors); i++ {
		if(floors[i].Name == name) {
			render(c, gin.H{
				"title": "Map",
				"payload": floors,
				"Image": "../" + floors[i].ImageFile,
				}, "index.html")
			println(floors[i].ImageFile)
		}
	}
}