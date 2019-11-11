package model

import (
	"CASystem/pkg/infra/model/base"
	"time"
)

type LendingHistory struct {
	base.Model
	UserID      User
	EquipmentID Equipment
	LentAt      time.Time `json:"lent_at"`
	ReturnedAt  time.Time `json:"returned_at"`
}
