package model

import (
	"CASystem/pkg/infra/model/base"
	"time"
)

type User struct {
	base.Model
	FirebaseUID string
	AuthorityID string
	LeftOn      time.Time
}
