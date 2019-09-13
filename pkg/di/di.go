package di

import (
	"CASystem/domain/service"
)

// Demo is facr
var Demo service.Demo

// Init Cache, Di
func Init() {
	initDemo()
}

func initDemo() {
	// Demo = service.NewDemo()
}
