package testing

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	db "github.com/SCCapstone/BitCrunch/db"
)

func TestMain(m *testing.M) {
	fmt.Println("Running All Unit Tests...")
	fmt.Println("Making sure that we are in the correct directory...")
	err := os.Chdir("../../")
	if err != nil {
		// Something happened when trying to change dir!
		fmt.Println("chdir didn't work!", err)
		return
	}
	m.Run()
	// TODO run a check for any db files created, txt files in devices delete them
	fmt.Println("Cleaning up created files...")
	err = os.Remove("floors.db")
	if err != nil {
		fmt.Println("There's an error in finding and deleting this file!", err)
		// no return call
	}
	err = RemoveContents("devices")
	if err != nil {
		fmt.Println("An error was raised when cleaning the devices folder, ", err)
	}

	fmt.Println()
	fmt.Println("All Unit Tests Completed!")
}

// TODO: test Edit Device once recent pull is completed
func TestDevices(t *testing.T) {
	// create a device, run error check for creation
	// run create device x2, run get all devices, compare to manual addition
	// run delete device x2, run error check, check empty all devices to getalldevices
	ip := "127.0.0.1" // for funsies
	defaultImage := "static/assets/default_image"
	floorName := "testingFloor"
	_, err := db.CreateFloor(floorName, "PLACEHOLDERFORDEVICELIST") // we don't actually test the floor here!
	/* 	note: as the deviceList string is connected in such a way that only
	the floor uses the string, we are using a placeholder here
	This also allows for checking for errors using (space) in the device list
	*/
	if err != nil {
		t.Error("There was an error creating the floor! Please run TestFloors for more info!", err)
	}

	device1, err := db.CreateDevice("TestDevice1", ip, defaultImage, floorName)
	if err != nil {
		t.Error("CreateDevice has returned an error: ", err)
	}
	device2, err := db.CreateDevice("TestDevice2", ip, defaultImage, floorName)
	if err != nil {
		t.Error("CreateDevice has returned an error: ", err)
	}

	devices, err := db.GetAllDevicesForFloor(floorName)
	if err != nil {
		t.Error("GetAllDevicesForFloor has returned an error: ", err)
	}
	for _, device := range devices {
		if db.GetDeviceName(device) != db.GetDeviceName(device1) && db.GetDeviceName(device) != db.GetDeviceName(device2) {
			// get devices per floor returned a device not a part of the original duo
			t.Error("GetAllDevicesFloor returned a device not originally created. : ", device)
		}
	}
	err = db.DeleteDevice("TestDevice1", floorName)
	if err != nil {
		t.Error("DeleteDevice (straight string) has returned an error: ", err)
	}
	err = db.DeleteDevice((db.GetDeviceName(device2)), floorName)
	if err != nil {
		t.Error("DeleteDevice (db.getname) has returned an error: ", err)
	}
	devices, err = db.GetAllDevicesForFloor(floorName)
	if err != nil {
		t.Error("GetAllDevicesForFloor has returned an error: ", err)
	}
	if len(devices) != 0 {
		t.Error("DeleteDevice has not deleted every device. We are left with the following devices: ", devices)
	}
}
func TestFloors(t *testing.T) {
	// create a floor, run error check for creation
	// create floor, getallfloors, run error check
	// edit floor 2, check for errors -- possible usage for checkfloor here
	// run delete floor x2, run error check, check empty all floors to getallfloors
	ip := "127.0.0.1" // for funsies
	defaultImage := "static/assets/default_image"
	floorNameString := "testingFloor"
	secondaryFloorNameString := "testingFloor2"
	deviceList := "testingDevice"
	floor1, err := db.CreateFloor(floorNameString, deviceList)
	if err != nil {
		t.Error("CreateFloor has returned an error: ", err)
	}
	floor2, err := db.CreateFloor(secondaryFloorNameString, deviceList)
	if err != nil {
		t.Error("CreateFloor has returned an error for making a second floor: ", err)
	}
	floors, err := db.GetAllFloors()
	if err != nil {
		t.Error("GetAllFloors has returned an error: ", err)
	}
	for _, floor := range floors {
		if db.GetFloorName(floor) != db.GetFloorName(floor1) && db.GetFloorName(floor) != db.GetFloorName(floor2) {
			// get all floors returned a floor not a part of the original duo (order irrelevant)
			t.Error("GetAllFloors returned a floor not originally created. : ", floor)
		}
	}

	// checking edit floors, and placing the devices into the floor
	// both through normal creation and editing
	device1, err := db.CreateDevice("TestDevice1", ip, defaultImage, floorNameString)
	if err != nil {
		t.Error("Devices has returned an error! Please run TestDevices for more info!", err)
	}
	device2, err := db.CreateDevice("TestDevice2", ip, defaultImage, floorNameString)
	if err != nil {
		t.Error("Devices has returned an error! Please run TestDevices for more info!", err)
	}

	devices, err := db.GetAllDevicesForFloor(floorNameString)
	if err != nil {
		t.Error("GetAllDevicesForFloor has returned an error: ", err)
	}
	for _, device := range devices {
		if db.GetDeviceName(device) != db.GetDeviceName(device1) || db.GetDeviceName(device) != db.GetDeviceName(device2) {
			// get devices per floor returned a device not a part of the original duo
			t.Errorf("There was an error occurred when editing both devices into a floor!GetAllDevicesFloor returned a device not originally created!")
		}
	}
	err = db.DeleteDevice((db.GetDeviceName(device1)), floorNameString)
	if err != nil {
		t.Error("DeleteDevice has returned an error: ", err)
	}
	err = db.DeleteDevice((db.GetDeviceName(device2)), floorNameString)
	if err != nil {
		t.Error("DeleteDevice has returned an error: ", err)
	}
	devices, err = db.GetAllDevicesForFloor(floorNameString)
	if err != nil {
		t.Error("GetAllDevicesForFloor has returned an error: ", err)
	}
	if len(devices) != 0 {
		t.Error("DeleteDevice has not deleted every device. We are left with the following devices: ", devices)
	}
}

// helper function used to clean a directory
// https://stackoverflow.com/questions/33450980/how-to-remove-all-contents-of-a-directory-using-golang
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
