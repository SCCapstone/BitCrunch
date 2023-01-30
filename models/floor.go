// models.floor.go

package models

import (
	"errors"

	db "github.com/SCCapstone/BitCrunch/db"
)

/*
floor struct has a name and a file where its image is stored
*/
type Floor struct {
	Name      string `json:"Name"`
	ImageFile string `json:"Devices"`
}

/*
List of floors
*/
var floorList = []Floor{
	{Name: "Floor 1", ImageFile: "static/assets/floor1.png"},
	{Name: "Floor 2", ImageFile: "static/assets/floor2.png"},
}

/*
Return a list of all the floors
*/
func GetAllFloors() []Floor {
	return floorList
}

/*
Creates a new floor and adds it to the list
*/
func CreateNewFloor(name, file string) (*Floor, error) {
	if db.CheckFloor(name) == nil {
		return nil, errors.New("Floor name is not available")
	}
	f := Floor{Name: name, ImageFile: file}
	floorList = append(floorList, f)
	return &f, nil
}

/*
Looks for and returns specified floor
*/
func FindFloor(floorName string) *Floor {
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
func EditFloor(new_name, new_file, old_name string) (*Floor, error) {
	if db.CheckFloor(new_name) == nil {
		return nil, errors.New("Layer name is not available")
	}
	floor_to_edit := FindFloor(old_name)
	floor_to_edit.Name = new_name
	floor_to_edit.ImageFile = new_file
	return floor_to_edit, nil
}
