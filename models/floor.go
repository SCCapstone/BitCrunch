// models.floor.go

package models

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
	{Name: "Floor 1", ImageFile: "static/assets/floor1.png"},
	{Name: "Floor 2", ImageFile: "static/assets/floor2.png"},
}

/*
Return a list of all the floors
*/
func GetAllFloors() []floor {
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
