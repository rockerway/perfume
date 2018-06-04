package models

import (
	"perfume/packages/database"
	"strconv"
)

// Participant struct
type Participant struct {
	*database.Model
}

// AddParticipant is the method that add participant
func (participant *Participant) AddParticipant(activityID int, userID int) {
	defer participant.Init().Close()
	participant.Exec("INSERT INTO participants (`activity_id`, `user_id`) VALUES (" + strconv.Itoa(activityID) + "," + strconv.Itoa(userID) + ");")
}

// DeleteParticipantByActivityIDAndUserID is the method that delete a participant by activity id and user id
func (participant *Participant) DeleteParticipantByActivityIDAndUserID(activityID string, userID int) {
	defer participant.Init().Close()
	participant.Query("DELETE FROM participants WHERE activity_id=" + activityID + " AND user_id=" + strconv.Itoa(userID))
}

// DeleteParticipantsByActivityID is the method that delete participants by activity id
func (participant *Participant) DeleteParticipantsByActivityID(activityID string) {
	defer participant.Init().Close()
	participant.Query("DELETE FROM participants WHERE activity_id=" + activityID)
}

// InitParticipant to init participant
func InitParticipant() *Participant {
	return &Participant{Model: database.InitModel()}
}
