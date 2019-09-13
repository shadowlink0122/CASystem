package main

import (
	"net/http/cgi"

	"github.com/shadowlink0122/CASystem/interfaces"
)

func main() {
	// demo
	interfaces.Serve()

	// Run CGI Server
	cgi.Serve(nil)
}
