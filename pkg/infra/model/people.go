package model

import (
	"CASystem/pkg/infra/model/base"
	"time"
)

type People struct {
	base.Model
	UserID         User
	LastName       string    `json:"last_name"`
	FirstName      string    `json:"first_name"`
	LastNameKana   string    `json:"last_name_kana"`
	FirstNameKana  string    `json:"first_name_kana"`
	Gender         string    `json:"gender"`
	FacultyID      Faculty   `gorm:"foreignkey:FacultyRefer"`
	AdmittedOn     time.Time `json:"admitted_on"`
	StudentNumber  int       `json:"student_number"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	GroupID        Group     `gorm:"foreignkey:GroupRefer"`
	RoleID         Role      `gorm:"foreignkey:RoleRefer"`
	Interest       string    `json:"interest"`
	SpecialAbility string    `json:"special_ability"`
	Portfolio      string    `json:"portfolio"`
	ImageLink      string    `json:"image_link"`
}
