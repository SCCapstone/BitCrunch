package db

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
Adds a Floor to the database.
Returns an error if things
went wrong. nil otherwise.
*/
func (db *dbase) AddFloor(name, deviceFile string) error {
	// TODO
	return nil
}

/*
Ensures the name for a
floor has not already been used.
Returns nil if the name is good.
*/
func (db *dbase) checkFloor(name string) error {
	// TODO
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
