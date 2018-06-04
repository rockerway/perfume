package web

import (
	"perfume/packages/router"
)

// Open method to register router url
func Open() {
	// router.RegisterByHandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	// 	var tpl = template.Must(template.ParseGlob("resources/*.gohtml"))
	// 	err := tpl.ExecuteTemplate(res, "home.gohtml", nil)
	// 	if !recoder.Write(err) {
	// 		io.WriteString(res, err.Error())
	// 	}
	// })
	router.RegisterByString("/", "HomeController", "Index")
	router.RegisterByString("/login", "LoginController", "Index")
	router.RegisterByString("/logincheck", "LoginController", "LoginCheck")
	router.RegisterByString("/logout", "LoginController", "Logout")
	router.RegisterByString("/register", "RegisterController", "Index")
	router.RegisterByString("/registercheck", "RegisterController", "Register")
	router.RegisterByString("/activity", "ActivityController", "Index")
	router.RegisterByString("/activity/create", "ActivityController", "Create")
	router.RegisterByString("/activity/store", "ActivityController", "Store")
	router.RegisterByString("/activity/delete", "ActivityController", "Delete")
	router.RegisterByString("/activity/join", "ActivityController", "Join")
	router.Start()
}
