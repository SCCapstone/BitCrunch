// models.floor.go

package models

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
	f := Floor{Name: name, ImageFile: file}
	floorList = append(floorList, f)
	return &f, nil
}
