package models

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

type ErrMessage struct {
	Message string
	ErrNum  int
	Time    time.Time
}

func (e *ErrMessage) String() string {
	date := e.Time
	dateString := date.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("ошибка: %s, ее время: %s, номер ошибки: %v", e.Message, dateString, e.ErrNum)
}

type NotificationMessage struct {
	Pair     string
	Price    decimal.Decimal
	OldPrice decimal.Decimal
	Time     time.Time
}

func (n *NotificationMessage) String() string {
	date := n.Time
	dateString := date.Format("2006-01-02 15:04:05")
	price := n.Price.String()
	oldPrice := n.OldPrice.String()
	return fmt.Sprintf("Пара %s\n теперь стоит %s\n старая цена %s\n время %s\n", n.Pair, price, oldPrice, dateString)
}
