package db

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const devices = "devices.db"

type device struct {
	name string
	ip   string // static IP of the device
	// This is the NAME of the image file
	// NOT the actual image in memory
	image string
	// This is the name of the floor that the device
	// should be attached to
	floorName string
}

/*
Add a device to the database.
Returns error if things went wrong.
nil otherwise
*/
func CreateDevice(name, ip, image, floorNm string) (dev device, err error) {
	// Check device name
	// if err = CheckDevice(name); err != nil {
	// 	return
	// }
	// Check IP formatting
	if err = CheckIP(ip); err != nil {
		return
	}
	// Making sure the floor can be read or
	// that it exists
	if _, err = ReadFloor(floorNm); err != nil {
		return
	}
	// All checks are good, creating the
	// floor and writing to db
	dev.name = name
	dev.ip = ip
	dev.image = image
	dev.floorName = floorNm

	// writing to db, might have errors
	if err = writeDevice(dev); err != nil {
		return device{}, err
	}
	// all is good
	return
}

/*
Writes a Device to the device db.
Returns and error if things didn't
go well. nil otherwise.
Should only be used in the
CreateDevice function.
*/
func writeDevice(d device) (err error) {
	fil, err := os.OpenFile(d.floorName + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fil.Close()

	// Creating the string from the device details
	// will append to the file
	writeString := fmt.Sprintf("%s\t%s\t%s\t%s\n", d.name, d.ip, d.image, d.floorName)
	_, err = fil.WriteString(writeString)
	if err != nil {
		return err
	}
	// All is good
	return nil
}

/*
This checks that the string fits
a valid IP format. Returns
nil if it's good. Error otherwise.
*/
func CheckIP(ip string) error {
	checkIP := strings.Split(ip, ".")
	if len(checkIP) != 4 {
		return fmt.Errorf("Incorrect format!")
	}
	// Getting each octet to verify that it is in
	// a valid IP range of 0-255
	first, err := strconv.Atoi(checkIP[0])
	if err != nil {
		return fmt.Errorf("First octet is bad!")
	}
	second, err := strconv.Atoi(checkIP[1])
	if err != nil {
		return fmt.Errorf("Second octet is bad!")
	}
	third, err := strconv.Atoi(checkIP[2])
	if err != nil {
		return fmt.Errorf("Third octet is bad!")
	}
	fourth, err := strconv.Atoi(checkIP[3])
	if err != nil {
		return fmt.Errorf("Fourth octet is bad!")
	}

	// Actual checking
	if first < 0 || first > 255 {
		return fmt.Errorf("First octet is bad!")
	}
	if second < 0 || second > 255 {
		return fmt.Errorf("Second octet is bad!")
	}
	if third < 0 || second > 255 {
		return fmt.Errorf("Third ocetet is bad!")
	}
	if fourth < 0 || second > 255 {
		return fmt.Errorf("Fourth octet is bad!")
	}

	// All should be good
	return nil
}

/*
This checks to ensure that no other
device has the same name in the db.
Return nil if good, error otherwise.
// */
// func CheckDevice(name string) error {
// 	fmt.Println("here3")
// 	_, err := ReadDevice(name)
// 	if err != nil {
// 		// No errors found
// 		// Which means the device was found
// 		return fmt.Errorf("Device name already in use!")
// 	}
// 	fmt.Println("here4")
	// return nil
// }

/*
Remove a device from the database.
Returns nil if it was sucessful.
Error otherwise.
*/
func DeleteDevice(name string) error {
	// making temp file
	delMe, err := os.Create(fmt.Sprintf("temp%s.tmp", name))
	if err != nil {
		return err
	}
	// Opening db file
	fi, err := os.Open(devices)
	if err != nil {
		return err
	}
	// Reading file to find the device to be removed
	scan := bufio.NewScanner(fi)
	var line string
	for scan.Scan() {
		line = scan.Text()
		if strings.Split(line, "\t")[0] != name {
			delMe.WriteString(line)
		}
	}

	// Finished coping the device data to new file
	// Cleaning up
	fi.Close()
	err = os.Remove(devices)
	if err != nil {
		return err
	}

	// Renaming the new file withouth the
	// deleted device to the devices.db name
	err = os.Rename(delMe.Name(), devices)
	if err != nil {
		return err
	}

	// All good, clean up
	delMe.Close()
	return nil
}

/*
Returns the IP of a device
given a name.
Pretty useless function, but here it is.
*/
// func GetIP(name string) (string, error) {
// 	dev, err := ReadDevice(name)
// 	if err != nil {
// 		return "", fmt.Errorf("Device not found!")
// 	}
// 	return dev.ip, nil
// }

/*
This function can be used to get every
device attached to a certain floor.
Returns a slice of devices and an error.
The slice will be empty if no devices are found
and the error will be nil.
The only possible non-nil error is if there is a problem
reading the devices.db file.
*/
func GetAllDevicesForFloor(floorNm string) (devs []device, err error) {
	fi, err := os.Open(floorNm + ".txt")
	if err != nil {
		return
	}
	defer fi.Close()
	scan := bufio.NewScanner(fi)
	var line []string
	// Finding each device with the given floor name
	firstLine := true
	for scan.Scan() {
	if firstLine == true {
		firstLine = false
		continue
	}
		line = strings.Split(scan.Text(), "\t")
			// Device found, append it to the slice
			d := device{
				name:      line[0],
				ip:        line[1],
				image:     line[2],
				floorName: line[3],
			}
			devs = append(devs, d)
	}

	return devs, nil
}
