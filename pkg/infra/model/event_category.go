package model

import "CASystem/pkg/infra/model/base"

type EventCategory struct {
	base.Model
	Name string `json:"name"`
}
