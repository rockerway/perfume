package router

import (
	"perfume/controllers"
)

// ClassMap is the global class map
var ClassMap map[string]interface{}

func registerClass(name string, class interface{}) {
	ClassMap[name] = class
}

func init() {
	registerClass("Controller", controllers.InitController())
	registerClass("LoginController", controllers.InitLoginController(ClassMap["Controller"]))

}
