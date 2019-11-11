package model

import "CASystem/pkg/infra/model/base"

type Project struct {
	base.Model
	Name            string        `json:"name"`
	About           string        `json:"about"`
	Comment         string        `json:"comment"`
	ProposalLink    string        `json:"proposal_link"`
	ProjectStatusID ProjectStatus `gorm:"foreignkey:ProjectStatusRefer"`
}
