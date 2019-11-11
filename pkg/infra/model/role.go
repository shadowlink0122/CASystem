package model

import "CASystem/pkg/infra/model/base"

type Role struct {
	base.Model
	Name string `json:"name"`
}
