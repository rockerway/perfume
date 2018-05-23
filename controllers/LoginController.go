package controllers

import (
	"net/http"
	"perfume/models"
)

// LoginController handle login function
type LoginController struct {
	*Controller
	user models.User
}

// Index is the browse view
func (controller *LoginController) Index(res http.ResponseWriter, req *http.Request) {
	if controller.Controller.auth.IsLogin(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	controller.render(res, "login.gohtml", nil)
}

// LoginCheck is the method that auth accoute
func (controller *LoginController) LoginCheck(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if controller.Controller.auth.IsLogin(res, req) {
			http.Redirect(res, req, "/", http.StatusSeeOther)
			return
		}
		email := req.FormValue("input_email")
		password := req.FormValue("input_password")
		firstName, lastName, hashedPassword := controller.user.GetUserByEmail(email)
		if controller.Controller.auth.Check(hashedPassword, password) {
			controller.Controller.auth.CreateSession(firstName, lastName, email, res)
			http.Redirect(res, req, "/", http.StatusSeeOther)
		} else {
			http.Error(res, "Username and/or password do not match", http.StatusForbidden)
		}
		return
	}
	http.Error(res, "No page. Did U understand???", http.StatusNotFound)
}

// Logout is the method to logout
func (controller *LoginController) Logout(res http.ResponseWriter, req *http.Request) {
	controller.Controller.auth.ClearSession(res, req)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}

// InitLoginController to init login controller
func InitLoginController(controller *Controller) *LoginController {
	return &LoginController{Controller: controller, user: models.InitUser()}
}
