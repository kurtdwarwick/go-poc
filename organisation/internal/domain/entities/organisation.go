package entities

import "shared"

type Organisation struct {
	shared.Entity
	LegalName   string
	TradingName string
	Website     *string
}
