package model

import (
	"CASystem/pkg/infra/model/base"
	"time"
)

type LendingHistory struct {
	base.Model
	UserID      User      `gorm:"foreignkey:UserRefer"`
	EquipmentID Equipment `gorm:"foreignkey:EquipmentRefer"`
	LentAt      time.Time `json:"lent_at"`
	ReturnedAt  time.Time `json:"returned_at"`
}
