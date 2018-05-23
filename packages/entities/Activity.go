package entities

// Activity entity
type Activity struct {
	oldDrive   Person
	passengers []Person
	date       Date
}

// GetOldDrive is the method that get some person
func (a Activity) GetOldDrive() Person {
	return a.oldDrive
}

// AddPassenger is the method that add a person to list
func (a Activity) AddPassenger(person Person) {
	a.passengers = append(a.passengers, person)
}

// GetPassengers is the method that get all passengers
func (a Activity) GetPassengers() []Person {
	return a.passengers
}

// GetDate is the method that get activity date
func (a Activity) GetDate() Date {
	return a.date
}

// InitActivity is the method that initial struct activity
func InitActivity(oldDrive Person, date Date) Activity {
	activity := Activity{
		oldDrive: oldDrive,
		date:     date,
	}
	return activity
}
