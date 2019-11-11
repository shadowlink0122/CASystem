package model

import "CASystem/pkg/infra/model/base"

type QuestionType struct {
	base.Model
	Name string `json:"name"`
}
