package models

import (
	"perfume/packages/database"
	"perfume/packages/exception"
	"strconv"
)

// Activity struct
type Activity struct {
	*database.Model
}

// GetAllActivities is the method that get all activities
func (activity *Activity) GetAllActivities(currentUserID int) []struct {
	ID          int
	Name        string
	Description string
	UserName    string
	IsCreator   bool
} {
	defer activity.Init().Close()
	result := []struct {
		ID          int
		Name        string
		Description string
		UserName    string
		IsCreator   bool
	}{}
	rows := activity.Query("select a.id, a.name, a.description, u.first_name, u.last_name, u.id from activities a right join users u on a.user_id=u.id")
	var userID, id int
	var name, description, userFirstName, userLastName string
	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &userFirstName, &userLastName, &userID)
		if recoder.Write(err) {
			result = append(result, struct {
				ID          int
				Name        string
				Description string
				UserName    string
				IsCreator   bool
			}{
				id,
				name,
				description,
				userFirstName + " " + userLastName,
				userID == currentUserID,
			})
		}
	}

	return result
}

// AddActivity is the method that add activity
func (activity *Activity) AddActivity(userID int, name, description string) {
	defer activity.Init().Close()
	activity.Query("INSERT INTO activities (`name`, `description`, `user_id`) VALUES (\"" + name + "\",\"" + description + "\"," + strconv.Itoa(userID) + ");")
}

// GetActivityByID is the method that get a activity by ID
func (activity *Activity) GetActivityByID(encodeID string) {
	defer activity.Init().Close()
}

// DeleteActivity is the method that delete a activity
func (activity *Activity) DeleteActivity(userID, activityID string) {
	defer activity.Init().Close()
	activity.Query("DELETE FROM activities WHERE id=" + activityID + " AND user_id=" + userID + ";")
}

// InitActivity to init activity
func InitActivity() *Activity {
	return &Activity{Model: database.InitModel()}
}
