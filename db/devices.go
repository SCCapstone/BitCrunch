package db

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const devices = "devices.db"

type device struct {
	name string
	ip   net.IP // static ip of device
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
	dev.ip = net.ParseIP(ip)
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
	writeString := fmt.Sprintf("%s\t%s\t%s\t%s\n", d.name, d.ip.String(), d.image, d.floorName)
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
func CheckIP(ip string) (net.IP, error) {
	// This function converts to a go IP type.
	// will be nil if parsing failed.
	check := net.ParseIP(ip)
	if check == nil {
		// Could be a hostname.
		// Getting the IP from net function
		hostIP, err := net.LookupHost(ip)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse IP address or host.")
		}
		// Just taking the first IP in the slice
		// Could potentially cause problems but
		// should work in most cases
		check = net.ParseIP(hostIP[0])
	}
	// IP is not nil so everything is good
	return check, nil
}

/*
This checks to ensure that no other
device has the same name in the db.
Return nil if good, error otherwise.
// */
// func CheckDevice(name string) error {
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
		d := device{
			name:      line[0],
			ip:        net.ParseIP(line[1]),
			image:     line[2],
			floorName: line[3],
		}
		devs = append(devs, d)
	}

	return devs, nil
}
