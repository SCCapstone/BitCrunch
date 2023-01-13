// models.floor.go

package main

/*
floor struct has a name and a file where its image is stored
*/
type floor struct {
	Name      string   `json:"Name"`
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
	f := floor{Name: name, ImageFile: file}
	floorList = append(floorList, f)
	return &f, nil
}