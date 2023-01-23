// handlers.map.go

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

/*
Renders the index with updated layer values
*/
func showMap(c *gin.Context) {
	floors := getAllFloors()

	render(c, gin.H{
		"title":   "Map",
		"payload": floors,
	}, "index.html")
}

/*
Adds a layer with a layer name inputted from the user
Saves uploaded image to static/assets folder
Creates a new floor and adds it to the list of floors, calls showMap to render the map with updates
*/
func addLayer(c *gin.Context) {
	layer_name := c.PostForm("layer_name")

	file, err := c.FormFile("layer_image")
	if err != nil {
		log.Println(err)
	}

	err = c.SaveUploadedFile(file, "static/assets/"+file.Filename)
	if err != nil {
		log.Println(err)
	}

	createNewFloor(layer_name, "static/assets/"+file.Filename)
	showMap(c)
}

/*
Changes layer name to new one inputted from the user
Saves uploaded image to static/assets folder and deletes previous image
*/
func editLayer(c *gin.Context) {
	new_layer_name := c.PostForm("layer_name")
	new_file, err := c.FormFile("layer_image")
	old_layer_name := c.PostForm("floor_name")
	if err != nil {
		log.Println(err)
	}
	remove := os.Remove(findFloor(old_layer_name).ImageFile)
	if remove != nil {
		log.Println(err)
	}
	err = c.SaveUploadedFile(new_file, "static/assets/"+new_file.Filename)
	editFloor(new_layer_name, "static/assets/"+new_file.Filename, old_layer_name)
	showMap(c)
}

/*
Gets the proper floor from the list of floors based on its name
Renders the proper floor image onto the map
*/
func viewLayer(c *gin.Context) {
	name := c.PostForm("l_name")
	floors := getAllFloors()
	for i := 0; i < len(floors); i++ {
		if floors[i].Name == name {
			render(c, gin.H{
				"title":   "Map",
				"payload": floors,
				"Image":   "../" + floors[i].ImageFile,
			}, "index.html")
		}
	}
}
