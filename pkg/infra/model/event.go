package model

import (
	"CASystem/pkg/infra/model/base"
	"time"
)

type Event struct {
	base.Model
	Name              string        `json:"name"`
	HeldOn            time.Time     `json:"held_on"`
	Summary           string        `json:"summary"`
	EventCategoriesID EventCategory `gorm:"foreignkey:EventCategoryRefer"`
}
