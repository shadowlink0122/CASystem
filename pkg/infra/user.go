package infra

import(
	"fmt"
)

type IUserRepo interface{
	UserList() ([]string, error)
}

type userRepo struct{
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}

func (urepo *userRepo) UserList() ([]string, error) {
	lcap := 10
	userList := make([]string, 0, lcap)

	for i := 0;i < lcap;i++{
		userName := fmt.Sprintf("Name%d", i)
		userList = append(userList, userName)
	}

	return userList, nil
}
