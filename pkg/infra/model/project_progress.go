package model

import "CASystem/pkg/infra/model/base"

type ProjectProgress struct {
	base.Model
	ProjectID Project `gorm:"foreignkey:ProjectRefer"`
	Progress  string  `json:"progress"`
}
