package model

import "CASystem/pkg/infra/model/base"

type ProjectStatus struct {
	base.Model
	Name string `json:"name"`
}
