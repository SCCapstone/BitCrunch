package db

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const floors = "floors.db"

type floor struct {
	name string
	/*
		The following will be a file
		which contains the names of
		each device for the floor.
	*/
	deviceListFile string
}

var floorList = []floor{}

// var currentFloor = ""
/*
Getter used for comparisons between devices in Unit Testing
*/
func GetFloorName(d floor) string {
	return d.name
}

/*
Creates a new floor.
Will check for valid name and file.
Returns error if things went wrong.
*/
func CreateFloor(name, deviceList string) (flo floor, err error) {
	flo = floor{
		name:           "",
		deviceListFile: "",
	}
	// Check floor name
	if err = CheckFloor(name); err != nil {
		return
	}
	// Check file name
	// if err = CheckFile(deviceList); err != nil {
	// 	return
	// }

	// Everything is good, so return the floor data
	flo.name = name
	flo.deviceListFile = deviceList

	if err = writeFloor(flo); err != nil {
		return floor{}, err
	}
	return
}

/*
Creates a new floor.
Will check for valid name and file.
Returns error if things went wrong.
*/
func EditFloor(name, deviceList string) (flo floor, err error) {
	flo = floor{
		name:           "",
		deviceListFile: "",
	}

	// Check floor name
	if err = CheckFloor(name); err != nil {
		return
	}
	// Check file name
	// if err = CheckFile(deviceList); err != nil {
	// 	return
	// }

	// Everything is good, so return the floor data
	flo.name = name
	flo.deviceListFile = deviceList

	if err = writeFloor(flo); err != nil {
		return floor{}, err
	}
	return
}

/*
Writes a Floor to the database.
Returns an error if things
went wrong. nil otherwise.
Should only be used by the
CreateFloor function
*/
func writeFloor(fl floor) error {
	fil, err := os.OpenFile(floors, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fil.Close()
	// Creating the string from the floor details
	// will append to the file
	writeString := fmt.Sprintf("%s\t%s\n", fl.name, fl.deviceListFile)
	_, err = fil.WriteString(writeString)
	if err != nil {
		return err
	}
	// Sucessful write
	return nil
}

func ReadFloor(fname string) (f floor, err error) {
	fi, err := os.Open(floors)
	if err != nil {
		return
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	for scan.Scan() {
		line = strings.Split(scan.Text(), "\t")
		//fmt.Println(line) commented out to reduce noise
		if line[0] == fname {
			f = floor{
				name:           line[0],
				deviceListFile: line[1],
			}
			return f, nil
		}
	}
	// The floor was not found
	// so return an error
	return floor{}, fmt.Errorf("Floor could not be read/found.")
}

/*
Ensures the name for a
floor has not already been used.
Returns nil if the name is good.
An error otherwise.
*/
func CheckFloor(name string) error {
	fi, err := open(floors)
	if err != nil {
		return err
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	for scan.Scan() {
		line = strings.Split(scan.Text(), "\t")
		if line[0] == name {
			return fmt.Errorf("Floor name \"%s\" is taken.", name)
		}
	}
	// The floor name was not found
	return nil
}

/*
Removes a floor from the database.
Returns nil if it was sucessful.
Returns error otherwise.
*/
func DeleteFloor(name string) error {
	// Creating a temp file
	delMe, err := os.Create(fmt.Sprintf("temp%s.tmp", name))
	newName := "floors.db"
	if err != nil {
		return err
	}
	fi, err := os.Open(floors)
	if err != nil {
		return err
	}
	scan := bufio.NewScanner(fi)
	var line string
	for scan.Scan() {
		line = scan.Text()
		if strings.Split(line, "\t")[0] != name {
			delMe.WriteString(line + "\n")
		}
	}
	// Done with the main file
	// Removing it
	fi.Close()
	err = os.Remove(floors)
	if err != nil {
		return err
	}

	// Renaming the file without the
	// floor to be deleted to the floors.db
	delMe.Close()
	err = os.Rename(delMe.Name(), newName)
	if err != nil {
		return err
	}

	// Done, clean up
	return nil
}

/*
Gets the file name of a floor
given a floor name.
Returns error if not sucessful.
Not a super useful function but it's
here anyway.
*/
func GetDeviceFile(floorName string) (string, error) {
	fl, err := ReadFloor(floorName)
	if err != nil {
		return "", err
	}
	return fl.deviceListFile, nil
}

/*
Returns a list of all the floors in the database (with file names)
*/
func GetAllFloors() (myfloors []floor, err error) {
	var floorList = []floor{}
	fi, err := os.Open(floors)
	if err != nil {
		return
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	for scan.Scan() {
		line = strings.Split(scan.Text(), "\t")
		f := floor{
			name:           line[0],
			deviceListFile: line[1],
		}
		floorList = append(floorList, f)
	}
	return floorList, nil
}
