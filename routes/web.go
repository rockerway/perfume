package web

import (
	"perfume/controllers"
	"perfume/packages/router"
)

// Open method to register router url
func Open() {
	webrouter.Register("/login", controllers.AuthorizeController{}.Index)
	webrouter.Register("/activity", controllers.ActivityController{}.Index)
	webrouter.Start()
}
