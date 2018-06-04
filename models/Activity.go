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
func (activity *Activity) GetAllActivities(currentUserID int, isLogin bool) []struct {
	ID            int
	Name          string
	Description   string
	UserName      string
	IsCreator     bool
	IsParticipant bool
	IsLogin       bool
} {
	defer activity.Init().Close()
	result := []struct {
		ID            int
		Name          string
		Description   string
		UserName      string
		IsCreator     bool
		IsParticipant bool
		IsLogin       bool
	}{}
	rows := activity.Query("select a.id, a.name, a.description,a.user_id, u.first_name, u.last_name, p.user_id from activities a join users u on a.user_id=u.id left join participants p on a.id=p.activity_id && p.user_id=" + strconv.Itoa(currentUserID))
	var participantID interface{}
	var userID, id int
	var name, description, userFirstName, userLastName string
	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &userID, &userFirstName, &userLastName, &participantID)
		if recoder.Write(err) {
			result = append(result, struct {
				ID            int
				Name          string
				Description   string
				UserName      string
				IsCreator     bool
				IsParticipant bool
				IsLogin       bool
			}{
				id,
				name,
				description,
				userFirstName + " " + userLastName,
				userID == currentUserID,
				participantID != nil,
				isLogin,
			})
		}
	}

	return result
}

// AddActivity is the method that add activity
func (activity *Activity) AddActivity(userID int, name, description string) int64 {
	defer activity.Init().Close()
	return activity.Exec("INSERT INTO activities (`name`, `description`, `user_id`) VALUES (\"" + name + "\",\"" + description + "\"," + strconv.Itoa(userID) + ");")
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
