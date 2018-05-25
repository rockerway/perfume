package controllers

import (
	"io"
	"net/http"
	"perfume/packages/auth"
	"perfume/packages/exception"
)

// ActivityController handle auth function
type ActivityController struct {
	*Controller
}

// Index is the browse view
func (controller *ActivityController) Index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "activity")
	recoder.Write(err)
}

// InitActivityController to initial activity controller
func InitActivityController(auth *auth.Authentication) *ActivityController {
	return &ActivityController{
		Controller: InitController(auth),
	}
}
