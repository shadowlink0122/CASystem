package model

import "CASystem/pkg/infra/model/base"

type UsersEvent struct {
	base.Model
	UserID     User  `gorm:"foreignkey:UserRefer"`
	EventID    Event `gorm:"foreignkey:EventRefer"`
	IsAttended bool  `json:"id_attended"`
	Amount     int   `json:"money"`
}
