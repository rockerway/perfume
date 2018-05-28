package controllers

import (
	"net/http"
	"perfume/packages/auth"
)

type HomeController struct {
	*Controller
}

// Index is the browse view
func (controller *HomeController) Index(res http.ResponseWriter, req *http.Request) {
	controller.render(res, req, "home.gohtml", nil, 1)
}

func InitHomeController(auth *auth.Authentication) *HomeController {
	return &HomeController{Controller: InitController(auth)}
}
