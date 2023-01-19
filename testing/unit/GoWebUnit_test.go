package testing

import (
	"fmt"
	"testing"

	models "github.com/SCCapstone/BitCrunch/models"
)

func TestMain(m *testing.M) {
	m.Run()
	fmt.Println("placeholder")
}

func Test1(t *testing.T) {
	//t.Errorf("This is a testing error in the Test1!")

	StartFloorList := models.GetAllFloors()
	newFloorList := models.GetAllFloors()
	for index, element := range StartFloorList {
		if element != newFloorList[index] {
			// throw error here
			t.Errorf("GetAllFloors in modules.floor.go does not correctly add floors into the list of floors!")
		}
	}
	//t.Errorf("This is a testing error in the Test1!")

}
