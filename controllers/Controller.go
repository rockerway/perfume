package controllers

import (
	"html/template"
	"net/http"
	"perfume/packages/auth"
)

// ControllerInterface to define router sub element
type ControllerInterface interface {
	render(res http.ResponseWriter, req *http.Request, view string, data interface{})
	Index(res http.ResponseWriter, req *http.Request)
	Show(res http.ResponseWriter, req *http.Request)
	Create(res http.ResponseWriter, req *http.Request)
	Delete(res http.ResponseWriter, req *http.Request)
}

// Controller include all controllers basic method
type Controller struct {
	auth *auth.Authentication
}

func (controller *Controller) render(res http.ResponseWriter, req *http.Request, view string, data interface{}) {
	var tpl *template.Template = template.Must(template.ParseGlob("resources/*.gohtml"))
	templateData := struct {
		IsLogin bool
		Data    interface{}
	}{
		// controller.loginStatusToView(res, req),
		controller.auth.IsLogin(res, req),
		data,
	}
	tpl.ExecuteTemplate(res, view, templateData)
}

func (controller *Controller) Index(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Method Not Allowed", 405)
}

func (controller *Controller) Show(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Method Not Allowed", 405)
}

func (controller *Controller) Create(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Method Not Allowed", 405)
}

func (controller *Controller) Delete(res http.ResponseWriter, req *http.Request) {
	http.Error(res, "Method Not Allowed", 405)
}

func (controller *Controller) loginStatusToView(res http.ResponseWriter, req *http.Request) int {
	if controller.auth.IsLogin(res, req) {
		return 1
	}
	return 0
}

// InitController to initial controller
func InitController(auth *auth.Authentication) *Controller {
	return &Controller{auth: auth}
}
