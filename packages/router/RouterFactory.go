package router

import (
	"net/http"
	"perfume/config"
	"perfume/packages/exception"
	"reflect"
)

var router = make(map[string]func(http.ResponseWriter, *http.Request))

// Start method is web router's entry point.
func Start() {
	for path, method := range router {
		http.HandleFunc(path, method)
	}
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/resources/assets/", http.StripPrefix("/resources/assets", http.FileServer(http.Dir("resources/assets"))))
	http.Handle("/resources/images/", http.StripPrefix("/resources/images", http.FileServer(http.Dir("resources/images"))))
	err := http.ListenAndServe(":8080", nil)
	recoder.Write(err)
}

// RegisterByHandleFunc is the method to define route rule by HandleFunc
func RegisterByHandleFunc(path string, method func(res http.ResponseWriter, req *http.Request)) {
	router[path] = method
}

// RegisterByString is the method to define route rule by controller name and method name
func RegisterByString(path, controllerName, methodName string) {
	class := config.ClassMap[controllerName]
	controllerType := reflect.TypeOf(*class)
	controllerValue := reflect.ValueOf(*class)
	_, state := controllerType.MethodByName(methodName)
	if state {
		method := controllerValue.MethodByName(methodName)
		RegisterByHandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
			method.Call([]reflect.Value{reflect.ValueOf(res), reflect.ValueOf(req)})
		})
	}
}
