package cookies

import (
	"net/http"
)

// GetCookie is the method that get cookie frome request
func GetCookie(req *http.Request, name string) (*http.Cookie, bool) {
	cookie, err := req.Cookie(name)
	state := err == nil
	return cookie, state
}

// SetCookie is the method that set cookie on response
func SetCookie(res http.ResponseWriter, name string, value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:  name,
		Value: value,
	}
	http.SetCookie(res, cookie)
	return cookie
}
