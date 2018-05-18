package webrouter

import (
	"net/http"
	"perfume/packages/exception"
)

var router = make(map[string]func(http.ResponseWriter, *http.Request))

// Start method is web router's entry point.
func Start() {
	for path, method := range router {
		http.HandleFunc(path, method)
	}
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	recoder.Write(err)
}

// Register is the method to define route rule.
func Register(path string, method func(res http.ResponseWriter, req *http.Request)) {
	router[path] = method
}
