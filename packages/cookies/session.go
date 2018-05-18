package session

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// GetSession method is get session in client cookie.
func GetSession(res http.ResponseWriter, req *http.Request) string {
	session, err := req.Cookie("session")
	if err != nil {
		uuid, _ := uuid.NewV4()
		session = &http.Cookie{
			Name:  "session",
			Value: uuid.String(),
		}
		http.SetCookie(res, session)
	}
	return session.Value
}
