package model

import "CASystem/pkg/infra/model/base"

type QuietionnaireContent struct {
	base.Model
	QuestionnaireID Questionnaire
	Question        string `json:"question"`
	QuestionTypeID  QuestionType
	Comment         string `json:"comment"`
}
