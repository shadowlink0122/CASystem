package interfaces

import (
	"net/http"
	"net/http/cgi"

	"github.com/shadowlink0122/CASystem/domain/service"
	"github.com/shadowlink0122/CASystem/interfaces/di"
)

// Serve is
func Serve() {
	di.Init()

	// demo
	http.Handle("/demo", get(service.DemoMessage()))

	// Run CGI Server
	cgi.Serve(nil)
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}
