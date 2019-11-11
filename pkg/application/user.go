package application

import (
	"CASystem/pkg/infra"
)

type IUserApp interface {
	GetList() ([]string, error)
}

type userApp struct {
	infra.IUserRepo
}

func NewUserApp(repo infra.IUserRepo) IUserApp {
	return &userApp{repo}
}

func (uapp *userApp) GetList() ([]string, error) {
	user, err := uapp.UserList()
	if err != nil {
		return nil, err
	}

	return user, nil
}
