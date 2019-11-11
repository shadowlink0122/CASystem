package model

import "CASystem/pkg/infra/model/base"

type Authority struct {
	base.Model
	Name string `json:"name"`
}
