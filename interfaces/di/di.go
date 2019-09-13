package di

import (
	"github.com/shadowlink0122/CASystem/domain/service"
)

// Demo is facr
var Demo service.Demo

// Init Cache, Di
func Init() {
}

func initDemo() {
	demo = service.demo.New()
}
