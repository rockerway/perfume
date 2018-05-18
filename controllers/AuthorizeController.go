package controllers

import (
	"io"
	"net/http"
	"perfume/packages/exception"
)

// AuthorizeController handle auth function
type AuthorizeController struct {
	*Controller
}

// Index is the browse view
func (controller AuthorizeController) Index(res http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(res, "auth")
	recoder.Write(err)
}
