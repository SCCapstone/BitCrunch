package db

type device struct {
	name string
	ip   string // static IP of the device
}

func (db *dbase) AddDevice(name, ip string) {
	// TODO
}

/*
This checks that the string fits
a valid IP format. Returns
nil if it's good. Error otherwise.
*/
func checkIP(ip string) error {
	// TODO
	return nil
}

/*
This checks to ensure that no other
device has the same name in the db.
Return nil if good, error otherwise.
*/
func checkName(name string) error {
	// TODO
	return nil
}

/*
Remove a device from the database.
Returns nil if it was sucessful.
Error otherwise.
*/
func (db *dbase) DeleteDevice(name string) error {
	return nil
}

/*
Returns the IP of a device
given a name.
*/
func (db *dbase) GetIP(name string) string {
	// TODO
	return ""
}
