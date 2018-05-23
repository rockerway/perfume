package controllers

import (
	"html/template"
	"net/http"
	"perfume/packages/auth"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("resources/*.gohtml"))
}

// Controller include all controllers basic method
type Controller struct {
	auth auth.Authentication
}

func (controller *Controller) render(res http.ResponseWriter, view string, data interface{}) {
	tpl.ExecuteTemplate(res, view, data)
}

// InitController to initial controller
func InitController() *Controller {
	return &Controller{auth: *auth.InitAuthentication()}
}
