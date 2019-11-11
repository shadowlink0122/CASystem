package model

import (
	"CASystem/pkg/infra/model/base"
	"time"
)

type Budget struct {
	base.Model
	Summary  string    `json:"summary"`
	Amount   string    `json:"amount"`
	TradedOn time.Time `json:"traded_on"`
	Seller   string    `json:"seller"`
	Comment  string    `json:"comment"`
}
