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
}
func TestFloorAddDelete(t *testing.T) {
	var floorList = []models.Floor{
		{Name: "Floor 1", ImageFile: "static/assets/floor1.png"},
		{Name: "Floor 2", ImageFile: "static/assets/floor2.png"},
	} // Taken directly from floor.go's default floors
	floorList = append(floorList, models.Floor{Name: "New Floor", ImageFile: "static/assets/floor1.png"})
	models.CreateNewFloor("New Floor", "static/assets/floor1.png")
	newFloorList := models.GetAllFloors() // GetAllFloors uses one slice for all floors, and is kept as "sudo-static"
	for index, element := range floorList {
		if element != newFloorList[index] {
			// throw error here
			t.Errorf("CreateNewFloor in modules.floor.go does not correctly add a new floor!")
		}
	}
	/*floorList = **remove from testing list
	** call floor removal
	for index, element := range floorList {
		if element != newFloorList[index] {
			// throw error here
			t.Errorf("DeleteFloor in modules.floor.go does not correctly remove a floor!")
		}
	}
	this is done because the floor database is static (to my understanding)
	*/
}

func TestUserValid(t *testing.T) {
	var userList = []models.User{
		{Username: "user1", Password: "pass1"},
		{Username: "user2", Password: "pass2"},
		{Username: "user3", Password: "pass3"},
	} // Taken directly from user.go's default users
	// going to wait until everything is in main to do these bad boys
	t.Errorf("These tests aren't ready yet! \n %s", userList)
}
func TestUserRegistration(t *testing.T) {
	var userList = []models.User{
		{Username: "user1", Password: "pass1"},
		{Username: "user2", Password: "pass2"},
		{Username: "user3", Password: "pass3"},
	} // Taken directly from user.go's default users
	// going to wait until everything is in main to do these bad boys
	t.Errorf("These tests aren't ready yet! \n %s", userList)
}
func TestUserAvaliable(t *testing.T) {
	var userList = []models.User{
		{Username: "user1", Password: "pass1"},
		{Username: "user2", Password: "pass2"},
		{Username: "user3", Password: "pass3"},
	} // Taken directly from user.go's default users
	// going to wait until everything is in main to do these bad boys
	t.Errorf("These tests aren't ready yet! \n %s", userList)
}