package config

import (
	"perfume/controllers"
	"perfume/packages/auth"
)

// ClassMap is the global class map
var ClassMap map[string]*controllers.ControllerInterface = make(map[string]*controllers.ControllerInterface)

func registerClass(name string, class controllers.ControllerInterface) {
	ClassMap[name] = &class
}

func init() {
	var auth *auth.Authentication = auth.InitAuthentication()
	registerClass("HomeController", controllers.InitHomeController(auth))
	registerClass("LoginController", controllers.InitLoginController(auth))
	registerClass("RegisterController", controllers.InitRegisterController(auth))
	registerClass("ActivityController", controllers.InitActivityController(auth))
}
