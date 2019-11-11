package model

import "CASystem/pkg/infra/model/base"

type UserProject struct {
	base.Model
	UserID    User
	ProjectID Project
}
