package model

import "CASystem/pkg/infra/model/base"

type ProjectProgress struct {
	base.Model
	ProjectID Project
	Progress  string `json:"progress"`
}
