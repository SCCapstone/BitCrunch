// models.floor.go

package main

type floor struct {
	Name      string   `json:"Name"`
	ImageFile string `json:"Devices"`
}

var floorList = []floor{
	floor{Name: "Floor 1", ImageFile: "static/assets/floor1.png"},
	floor{Name: "Floor 2", ImageFile: "static/assets/floor2.png"},
}

// Return a list of all the floors
func getAllFloors() []floor {
	return floorList
}

func createNewFloor(name, file string) (*floor, error) {
	// Set the ID of a new article to one more than the number of articles
	f := floor{Name: name, ImageFile: file}
	floorList = append(floorList, f)

	return &f, nil
}