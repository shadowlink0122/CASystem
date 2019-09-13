package main

import (
	"net/http/cgi"

	"CASystem/pkg"
)

func main() {
	// demo
	pkg.Serve()

	// Run CGI Server
	cgi.Serve(nil)
}
