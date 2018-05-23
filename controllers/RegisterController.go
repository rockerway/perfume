package controllers

import (
	"net/http"
	"perfume/models"
)

// RegisterController handle register function
type RegisterController struct {
	*Controller
	user models.User
}

// Index is the browse view
func (controller *RegisterController) Index(res http.ResponseWriter, req *http.Request) {
	controller.user.GetUsers()
	if controller.Controller.auth.IsLogin(res, req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	controller.render(res, "register.gohtml", nil)
}

// Register is the method that add new user
func (controller *RegisterController) Register(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		firstName := req.FormValue("text_first_name")
		lastName := req.FormValue("text_last_name")
		email := req.FormValue("text_email")
		password := req.FormValue("text_password")
		if controller.isUserExist(email) {
			http.Error(res, "Email already taken", http.StatusForbidden)
			return
		}
		controller.user.AddUser(firstName, lastName, email, password)
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	http.Error(res, "No page. Did U understand???", http.StatusNotFound)
}

func (controller *RegisterController) isUserExist(email string) bool {
	firstName, lastName, password := controller.user.GetUserByEmail(email)
	if firstName == "" && lastName == "" && password == "" {
		return false
	}
	return true
}

// InitRegisterController to init register controller
func InitRegisterController() *RegisterController {
	return &RegisterController{
		Controller: InitController(),
		user:       models.InitUser(),
	}
}
