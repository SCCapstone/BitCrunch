package testing

import (
	"fmt"
	"testing"

	models "github.com/SCCapstone/BitCrunch/models"
)

func TestMain(m *testing.M) {
	fmt.Println("Running All Tests...")
	m.Run()
	fmt.Println("All Unit Tests Completed!")
}

func TestFloorModel(t *testing.T) {
	var floorList = []models.Floor{
		{Name: "Floor 1", ImageFile: "static/assets/floor1.png"},
		{Name: "Floor 2", ImageFile: "static/assets/floor2.png"},
	} // Taken directly from floor.go's default floors
	newFloorList := models.GetAllFloors()
	for index, element := range floorList {
		if element != newFloorList[index] {
			// throw error here
			t.Errorf("GetAllFloors in modules.floor.go does not correctly display the current floors!")
		}
	}
	floorList = append(floorList, models.Floor{Name: "New Floor", ImageFile: "static/assets/floor1.png"})
	models.CreateNewFloor("New Floor", "static/assets/floor1.png")
	newFloorList = models.GetAllFloors() // GetAllFloors uses one slice for all floors, and is kept as "sudo-static"
	for index, element := range floorList {
		if element != newFloorList[index] {
			// throw error here
			t.Errorf("CreateNewFloor in modules.floor.go does not correctly display add a new floors!")
		}
	}
	//t.Errorf("This is a testing error in the Test1!")

}
