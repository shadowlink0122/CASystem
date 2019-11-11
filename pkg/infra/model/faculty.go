package model

import "CASystem/pkg/infra/model/base"

type Faculty struct {
	base.Model
	Name  string `json:"name"`
	About string `json:"about"`
}
