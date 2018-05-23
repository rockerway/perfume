package web

import (
	"html/template"
	"io"
	"net/http"
	"perfume/controllers"
	"perfume/packages/exception"
	"perfume/packages/router"
)

// Open method to register router url
func Open() {
	webrouter.Register("/", func(res http.ResponseWriter, req *http.Request) {
		var tpl = template.Must(template.ParseGlob("resources/*.gohtml"))
		err := tpl.ExecuteTemplate(res, "home.gohtml", nil)
		if !recoder.Write(err) {
			io.WriteString(res, err.Error())
		}
	})
	webrouter.Register("/login", controllers.InitLoginController().Index)
	webrouter.Register("/logincheck", controllers.InitLoginController().LoginCheck)
	webrouter.Register("/register", controllers.InitRegisterController().Index)
	webrouter.Register("/registercheck", controllers.InitRegisterController().Register)
	webrouter.Register("/activity", controllers.InitActivityController().Index)
	webrouter.Start()
}
