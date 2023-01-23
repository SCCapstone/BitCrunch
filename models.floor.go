// models.floor.go

package main

import (
	"errors"
)

/*
floor struct has a name and a file where its image is stored
*/
type floor struct {
	Name      string `json:"Name"`
	ImageFile string `json:"Devices"`
}

/*
List of floors
*/
var floorList = []floor{
	floor{Name: "Floor 1", ImageFile: "static/assets/floor1.png"},
	floor{Name: "Floor 2", ImageFile: "static/assets/floor2.png"},
}

/*
Return a list of all the floors
*/
func getAllFloors() []floor {
	return floorList
}

/*
Creates a new floor and adds it to the list
*/
func createNewFloor(name, file string) (*floor, error) {
	if !isLayerNameAvailable((name)) {
		return nil, errors.New("Layer name is not available")
	}

	f := floor{Name: name, ImageFile: file}
	floorList = append(floorList, f)
	return &f, nil
}

/*
Looks for and returns specified floor
*/
func findFloor(floorName string) *floor {
	for i := 0; i < len(floorList); i++ {
		if floorList[i].Name == floorName {
			return &floorList[i]
		}
	}
	return nil
}

/*
Changes name and file of currently existing floor
*/
func editFloor(new_name, new_file, old_name string) (*floor, error) {
	if !isLayerNameAvailable(new_name) {
		return nil, errors.New("Layer name is not available")
	}
	floor_to_edit := findFloor(old_name)
	floor_to_edit.Name = new_name
	floor_to_edit.ImageFile = new_file
	return floor_to_edit, nil
}

/*
Checks if layer name is available
*/
func isLayerNameAvailable(name string) bool {
	for _, f := range floorList {
		if f.Name == name {
			return false
		}
	}
	return true
}
