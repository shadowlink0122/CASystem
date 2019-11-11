package model

import "CASystem/pkg/infra/model/base"

type Questionnaire struct {
	base.Model
	Title   string `json:"title"`
	Comment string `json:"comment"`
}
