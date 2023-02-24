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
}

var floorList = []floor{}

/*
Creates a new floor.
Will check for valid name and file.
Returns error if things went wrong.
*/
func CreateFloor(name string) (flo floor, err error) {
	flo = floor{
		name: "",
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
	writeString := fmt.Sprintf("%s\t \n", fl.name)
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
		if line[0] == fname {
			f = floor{
				name: line[0],
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
	fi, err := os.Open(floors)
	if err != nil {
		return err
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	for scan.Scan() {
		line = strings.Split(scan.Text(), "\t")
		if line[0] == name {
			return fmt.Errorf("Floor name found!")
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
			delMe.WriteString(line)
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
	err = os.Rename(delMe.Name(), floors)
	if err != nil {
		return err
	}

	// Done, clean up
	delMe.Close()
	return nil
}

/*
Returns a list of all the floors in the database (with file names)
*/
func GetAllFloors() (myfloors []floor, err error) {
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
			name: line[0],
		}
		myfloors = append(myfloors, f)
	}
	return
}
