package model

import "CASystem/pkg/infra/model/base"

type UserProject struct {
	base.Model
	UserID    User    `gorm:"foreignkey:UserRefer"`
	ProjectID Project `gorm:"foreignkey:ProjectRefer"`
}
