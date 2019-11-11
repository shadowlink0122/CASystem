package model

import "CASystem/pkg/infra/model/base"

type Equipment struct {
	base.Model
	Name string `json:"name"`
}
