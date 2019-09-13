package main

import (
	"net/http/cgi"

	"CASystem/interfaces"
)

func main() {
	// demo
	interfaces.Serve()

	// Run CGI Server
	cgi.Serve(nil)
}
