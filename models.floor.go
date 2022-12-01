// models.floor.go

package main

type floor struct {
	Name      string   `json:"Name"`
	DeviceFile   string `json:"Devices"`
}

var floorList = []floor{
	floor{Name: "Floor 1", DeviceFile: "floor1.txt"},
	floor{Name: "Floor 2", DeviceFile: "floor2.txt"},
}

// Return a list of all the floors
func getAllFloors() []floor {
	return floorList
}

func createNewFloor(name, file string) (*floor, error) {
	// Set the ID of a new article to one more than the number of articles
	f := floor{Name: name, DeviceFile: file}

	floorList = append(floorList, f)

	return &f, nil
}