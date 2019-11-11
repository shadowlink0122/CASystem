package di

import (
	"CASystem/pkg/application"
	"CASystem/pkg/infra"
)

var User application.IUserApp

func Init() {
	initUser()
}

func initUser() {
	user := infra.NewUserRepo()
	User = application.NewUserApp(user)
}
