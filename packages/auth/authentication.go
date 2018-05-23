package auth

import (
	"net/http"
	"perfume/packages/cookies"
	"time"

	uuid "github.com/satori/go.uuid"
)

const sessionExistTime int = 150

var sessionsCleaned time.Time = time.Now()

type user struct {
	FirstName string
	LastName  string
}

type session struct {
	email        string
	lastActivity time.Time
}

// Authentication struct
type Authentication struct {
	loginUser   map[string]user
	userSession map[string]session
}

// IsLogin is the method that adject login state by session Id
func (authentication *Authentication) IsLogin(res http.ResponseWriter, req *http.Request) bool {
	sessionID, sessionIDState := cookies.GetCookie(req, "session")
	if !sessionIDState {
		return false
	}
	session, sessionState := authentication.userSession[sessionID.Value]
	if sessionState {
		session.lastActivity = time.Now()
		authentication.userSession[sessionID.Value] = session
	}
	_, userState := authentication.loginUser[session.email]
	sessionID.MaxAge = sessionExistTime
	http.SetCookie(res, sessionID)
	return userState
}

// Check is the method that login check
func (authentication *Authentication) Check(hashedPassword, password string) bool {
	return PasswordCheck(hashedPassword, password)
}

// CreateSession is the method that create session to keep user login
func (authentication *Authentication) CreateSession(firstName, lastName, email string, res http.ResponseWriter) {
	sessionID, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:  "session",
		Value: sessionID.String(),
	}
	cookie.MaxAge = sessionExistTime
	http.SetCookie(res, cookie)
	authentication.userSession[cookie.Value] = session{email, time.Now()}
	authentication.loginUser[email] = user{firstName, lastName}
}

// ClearSession is the method that clear timeout session
func (authentication *Authentication) ClearSession(res http.ResponseWriter, req *http.Request) {
	cookie, _ := cookies.GetCookie(req, "session")
	// delete the session
	delete(authentication.loginUser, authentication.userSession[cookie.Value].email)
	delete(authentication.userSession, cookie.Value)
	// remove the cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(res, cookie)

	// clean up dbSessions
	if time.Now().Sub(sessionsCleaned) > (time.Second * 30) {
		for sessionID, session := range authentication.userSession {
			if time.Now().Sub(session.lastActivity) > (time.Second * 30) {
				delete(authentication.loginUser, session.email)
				delete(authentication.userSession, sessionID)
			}
		}
		sessionsCleaned = time.Now()
	}
}

// InitAuthentication to initial authentication struct
func InitAuthentication() *Authentication {
	return &Authentication{
		map[string]user{},
		map[string]session{},
	}
}
