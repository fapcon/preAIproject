package models

import "github.com/shopspring/decimal"

type Signal struct {
	OrderType int
	Pair      string
	PairID    int
	PairPrice decimal.Decimal
	BotUUID   string
}
