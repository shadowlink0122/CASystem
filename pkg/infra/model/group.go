package model

import "CASystem/pkg/infra/model/base"

type Group struct {
	base.Model
	Name  string `json:"name"`
	About string `json:"about"`
}
