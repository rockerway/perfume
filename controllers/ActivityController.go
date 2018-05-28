package controllers

import (
	"net/http"
	"perfume/models"
	"perfume/packages/auth"
	"strconv"
)

// ActivityController handle auth function
type ActivityController struct {
	*Controller
	activity *models.Activity
}

// Index is the browse view
func (controller *ActivityController) Index(res http.ResponseWriter, req *http.Request) {
	user := controller.auth.GetUserData(res, req)
	activities := controller.activity.GetAllActivities(user.ID)
	controller.render(res, req, "activity.gohtml", activities, 1)
}

// Create is the create view
func (controller *ActivityController) Create(res http.ResponseWriter, req *http.Request) {
	if !controller.auth.IsLogin(res, req) {
		http.Redirect(res, req, "/login", http.StatusSeeOther)
		return
	}
	controller.render(res, req, "activity_create.gohtml", nil, 2)
}

// Store is the create action
func (controller *ActivityController) Store(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if !controller.auth.IsLogin(res, req) {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		name := req.FormValue("text_name")
		description := req.FormValue("text_description")
		user := controller.auth.GetUserData(res, req)
		controller.activity.AddActivity(user.ID, name, description)
		http.Redirect(res, req, "/activity", http.StatusSeeOther)
		return
	}
	http.Error(res, "No page. Did U understand???", http.StatusNotFound)
}

// Delete is the delete action
func (controller *ActivityController) Delete(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if !controller.auth.IsLogin(res, req) {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}
		user := controller.auth.GetUserData(res, req)
		activityID := req.FormValue("hidden_activity_id")
		controller.activity.DeleteActivity(strconv.Itoa(user.ID), activityID)
		http.Redirect(res, req, "/activity", http.StatusSeeOther)
		return
	}
	http.Error(res, "No page. Did U understand???", http.StatusNotFound)
}

// InitActivityController to initial activity controller
func InitActivityController(auth *auth.Authentication) *ActivityController {
	return &ActivityController{
		Controller: InitController(auth),
		activity:   models.InitActivity(),
	}
}
