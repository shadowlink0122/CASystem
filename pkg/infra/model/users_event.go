package model

import "CASystem/pkg/infra/model/base"

type UsersEvent struct {
	base.Model
	UserID     User
	EventID    Event
	IsAttended bool `json:"id_attended"`
	Amount     int  `json:"money"`
}
