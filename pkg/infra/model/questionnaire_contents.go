package model

import "CASystem/pkg/infra/model/base"

type QuietionnaireContent struct {
	base.Model
	QuestionnaireID Questionnaire `gorm:"foreignkey:QuiestionnaireRefer"`
	Question        string        `json:"question"`
	QuestionTypeID  QuestionType  `gorm:"foreignkey:QuestionTypeRefer"`
	Comment         string        `json:"comment"`
}
