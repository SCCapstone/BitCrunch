package db

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
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
Getter used for comparisons between devices in Unit Testing
*/
func GetDeviceName(d device) string {
	return d.name
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
		return dev, fmt.Errorf("IP format is invalid")
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
	fil, err := os.OpenFile("devices/"+d.floorName+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
device has the same name in the db for a specific floor.
Return nil if good, error otherwise.
*/

func CheckDevice(name, floorNm string) error {
	fi, err := os.Open("devices/" + floorNm + ".txt")
	if err != nil {
		return err
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
		if len(line) > 1 {
			d := device{
				name:      line[0],
				ip:        line[1],
				image:     line[2],
				floorName: line[3],
			}
			if name == d.name {
				return errors.New("Device name already exists for floor")
			}
		}
	}

	return nil
}

func EditDevice(name, newName, newIP, newImage, floorNm string) {
	fi, err := ioutil.ReadFile("devices/" + floorNm + ".txt")
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(fi), "\n")
	firstLine := true

	for i, line := range lines {
		if firstLine == true {
			firstLine = false
			continue
		}
		splitLine := strings.Split(line, "\t")
		fmt.Println(splitLine, len(splitLine))
		if len(splitLine) > 1 {
			d := device{
				name:      splitLine[0],
				ip:        splitLine[1],
				image:     splitLine[2],
				floorName: splitLine[3],
			}
			if d.name == name {
				writeString := fmt.Sprintf("%s\t%s\t%s\t%s", newName, newIP, newImage, d.floorName)
				lines[i] = writeString
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("devices/"+floorNm+".txt", []byte(output), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

/*
Remove a device from the database.
Returns nil if it was sucessful.
Error otherwise.
*/
func DeleteDevice(name, floorNm string) error {
	// making temp file
	delMe, err := os.Create(fmt.Sprintf("temp%s.tmp", name))
	newName := "devices/" + floorNm + ".txt"
	if err != nil {
		return err
	}
	// Opening db file
	fi, err := os.Open("devices/" + floorNm + ".txt")
	if err != nil {
		return err
	}
	// Reading file to find the device to be removed
	scan := bufio.NewScanner(fi)
	var line string
	for scan.Scan() {
		line = scan.Text()
		if strings.Split(line, "\t")[0] != name {
			delMe.WriteString(line + "\n")
		}
	}

	// Finished coping the device data to new file
	// Cleaning up
	fi.Close()
	err = os.Remove("devices/" + floorNm + ".txt")
	if err != nil {
		return err
	}

	// Renaming the new file withouth the
	// deleted device to the devices.db name
	delMe.Close()
	err = os.Rename(delMe.Name(), newName)
	if err != nil {
		return err
	}

	// All good, clean up
	return nil
}

func GetIP(name string) string {
	floors, err := GetAllFloors()
	if err != nil {
		fmt.Println(err)
	}
	for _, floor := range floors {
		devices, _ := GetAllDevicesForFloor(floor.name)
		for _, device := range devices {
			if device.name == name {
				return device.ip
			}
		}
	}
	return ""
}

func GetImage(name string) string {
	floors, err := GetAllFloors()
	if err != nil {
		fmt.Println(err)
	}
	for _, floor := range floors {
		devices, _ := GetAllDevicesForFloor(floor.name)
		for _, device := range devices {
			if device.name == name {
				return device.image
			}
		}
	}
	return ""
}

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
	fi, err := os.Open("devices/" + floorNm + ".txt")
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
		if len(line) > 1 {
			d := device{
				name:      line[0],
				ip:        line[1],
				image:     line[2],
				floorName: line[3],
			}
			devs = append(devs, d)
		}
	}

	return devs, nil
}

func GetAllIPs() (myDevices []string, err error) {
	var deviceIPs = []string{}
	floors, err := GetAllFloors()
	if err != nil {
		return deviceIPs, err
	}
	for _, floor := range floors {
		devices, _ := GetAllDevicesForFloor(floor.name)
		for _, device := range devices {
			deviceIPs = append(deviceIPs, device.ip)
		}
	}
	return deviceIPs, nil
}
