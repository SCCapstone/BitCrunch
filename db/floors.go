package db

import (
	"context"
	"fmt"
	"log"
	"time"
)

type floor struct {
	name string
	/*
		The following will be a file
		which contains the names of
		each device for the floor.
	*/
	deviceListFile string
}

/*
Creates a new floor.
Will check for valid name and file.
Returns error if things went wrong.
*/
func (db *base) New(name, deviceList string) (flo floor, err error) {
	flo = floor{
		name:           "",
		deviceListFile: "",
	}
	// Check floor name
	if err = db.checkFloor(name); err != nil {
		return
	}
	// Check file name
	if err = checkFile(deviceList); err != nil {
		return
	}

	// Everything is good, so return the floor data
	flo.name = name
	flo.deviceListFile = deviceList
	return
}

/*
Adds a Floor to the database.
Returns an error if things
went wrong. nil otherwise.
*/
func (db *dbase) AddFloor(fl floor) error {
	if !db.opened {
		db.Open()
	}

	query := "INSERT INTO floors(name, devicelist) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.sqldb.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, fl.name, fl.deviceListFile)
	if err != nil {
		return err
	}
	return nil
}

/*
Ensures the name for a
floor has not already been used.
Returns nil if the name is good.
*/
func (db *dbase) checkFloor(name string) error {
	if !db.opened {
		db.Open()
	}

	query := fmt.Sprintf("SELECT %s FROM floors", name)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.sqldb.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	var test string
	row := stmt.QueryRowContext(ctx, name)
	if err = row.Scan(&test); err == nil {
		return fmt.Errorf("Floor name exists!")
	}
	return nil
}

/*
Checks to ensure the device file
exists. Returns nil if it does.
*/
func checkFile(file string) error {
	// TODO
	return nil
}

/*
Removes a floor from the database.
Returns nil if it was sucessful.
Returns error otherwise.
*/
func (db *dbase) DeleteFloor(name string) error {
	// TODO
	return nil
}

/*
Gets the file name of a floor
given a floor name.
Returns error if not sucessful.
*/
func (db *dbase) GetDeviceFile(name string) (string, error) {
	// TODO
	return "", nil
}
