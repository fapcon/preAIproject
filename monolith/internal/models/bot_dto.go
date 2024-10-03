package models

import (
	"time"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/db/types"
)

const (
	BuyMarket = iota + 1
	SellMarket
)

var OrderTypes = map[int]string{
	BuyMarket: "BUYMARKET",
}

//go:generate easytags $GOFILE json,db,db_ops,db_type,db_default,mapper
type BotDTO struct { // DTO - data transfer object
	ID                   int               `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" db_ops:"id" mapper:"id"`
	Kind                 int               `json:"kind" db:"kind" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"kind"`
	UserID               int               `json:"user_id" db:"user_id" db_ops:"create" db_type:"int" db_default:"default 1" db_index:"index" mapper:"user_id"`
	Name                 string            `json:"name" db:"name" db_ops:"create,update" db_type:"varchar(55)" db_default:"not null" mapper:"name"`
	Description          string            `json:"description" db:"description" db_ops:"create,update" db_type:"varchar(144)" db_default:"not null" mapper:"description"`
	PairID               int               `json:"pair_id" db:"pair_id" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"pair_id"`
	FixedAmount          float64           `json:"fixed_amount" db:"fixed_amount" db_ops:"create,update" db_type:"decimal(21,8)" db_default:"default 0" mapper:"fixed_amount"`
	ExchangeType         int               `json:"exchange_type" db:"exchange_type" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"exchange_type"`
	ExchangeID           int               `json:"exchange_id" db:"exchange_id" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"exchange_id"`
	ExchangeUserKeyID    int               `json:"exchange_user_key_id" db:"exchange_user_key_id" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"exchange_user_key_id"`
	OrderType            int               `json:"order_type" db:"order_type" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"order_type"`
	SellPercent          float64           `json:"sell_percent" db:"sell_percent" db_ops:"create,update" db_type:"decimal(5,3)" db_default:"default 1" mapper:"sell_percent"`
	CommissionPercent    float64           `json:"commission_percent" db:"commission_percent" db_ops:"create,update" db_type:"decimal(5,3)" db_default:"default 13" mapper:"commission_percent"`
	AssetType            int               `json:"asset_type" db:"asset_type" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"asset_type"`
	UUID                 string            `json:"uuid" db:"uuid" db_ops:"create" db_type:"char(36)" db_default:"not null" db_index:"index,unique" mapper:"uuid"`
	Active               types.NullBool    `json:"active" db:"active" db_ops:"create,update" db_type:"boolean" db_default:"null" mapper:"active"`
	LimitOrder           types.NullBool    `json:"limit_order" db:"limit_order" db_ops:"create,update" db_type:"boolean" db_default:"null" mapper:"limit_order"`
	LimitSellPercent     types.NullFloat64 `json:"limit_sell_percent" db:"limit_sell_percent" db_ops:"create,update" db_type:"decimal(5,3)" db_default:"null" mapper:"limit_sell_percent"`
	LimitBuyPercent      types.NullFloat64 `json:"limit_buy_percent" db:"limit_buy_percent" db_ops:"create,update" db_type:"decimal(5,3)" db_default:"null" mapper:"limit_buy_percent"`
	AutoSell             types.NullBool    `json:"auto_sell" db:"auto_sell" db_ops:"create,update" db_type:"boolean" db_default:"null" mapper:"auto_sell"`
	AutoLimitSellPercent types.NullFloat64 `json:"auto_limit_sell_percent" db:"auto_limit_sell_percent" db_ops:"create,update" db_type:"decimal(5,3)" db_default:"null" mapper:"auto_limit_sell_percent"`
	OrderCountLimit      types.NullBool    `json:"order_count_limit" db:"order_count_limit" db_ops:"create,update" db_type:"boolean" db_default:"null" mapper:"order_count_limit"`
	OrderCount           int               `json:"order_count" db:"order_count" db_ops:"create,update" db_type:"int" db_default:"default 1" mapper:"order_count"`
	CreatedAt            time.Time         `json:"created_at" db:"created_at" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" db_ops:"created_at" mapper:"created_at"`
	UpdatedAt            time.Time         `json:"updated_at" db:"updated_at" db_ops:"update" db_type:"timestamp" db_default:"default (now()) not null" db_index:"index" mapper:"updated_at"`
	DeletedAt            types.NullTime    `json:"deleted_at" db:"deleted_at" db_ops:"update" db_type:"timestamp" db_default:"default null" db_index:"index" db_ops:"deleted_at" mapper:"deleted_at"`
}

func (s *BotDTO) TableName() string {
	return "bot"
}

func (s *BotDTO) OnCreate() []string {
	return []string{}
}

func (s *BotDTO) SetID(id int) *BotDTO {
	s.ID = id
	return s
}

func (s *BotDTO) GetID() int {
	return s.ID
}

func (s *BotDTO) SetKind(kind int) *BotDTO {
	s.Kind = kind
	return s
}

func (s *BotDTO) GetKind() int {
	return s.Kind
}

func (s *BotDTO) SetUserID(userId int) *BotDTO {
	s.UserID = userId
	return s
}

func (s *BotDTO) GetUserID() int {
	return s.UserID
}

func (s *BotDTO) SetName(name string) *BotDTO {
	s.Name = name
	return s
}

func (s *BotDTO) GetName() string {
	return s.Name
}

func (s *BotDTO) SetFixedAmount(fixedAmount float64) *BotDTO {
	s.FixedAmount = fixedAmount
	return s
}

func (s *BotDTO) GetFixedAmount() float64 {
	return s.FixedAmount
}

func (s *BotDTO) SetDescription(description string) *BotDTO {
	s.Description = description
	return s
}

func (s *BotDTO) GetDescription() string {
	return s.Description
}

func (s *BotDTO) SetPairID(pairId int) *BotDTO {
	s.PairID = pairId
	return s
}

func (s *BotDTO) GetPairID() int {
	return s.PairID
}

func (s *BotDTO) SetExchangeType(exchangeType int) *BotDTO {
	s.ExchangeType = exchangeType
	return s
}

func (s *BotDTO) GetExchangeType() int {
	return s.ExchangeType
}

func (s *BotDTO) SetExchangeID(exchangeId int) *BotDTO {
	s.ExchangeID = exchangeId
	return s
}

func (s *BotDTO) GetExchangeID() int {
	return s.ExchangeID
}

func (s *BotDTO) SetOrderType(orderType int) *BotDTO {
	s.OrderType = orderType
	return s
}

func (s *BotDTO) GetOrderType() int {
	return s.OrderType
}

func (s *BotDTO) SetSellPercent(sellPercent float64) *BotDTO {
	s.SellPercent = sellPercent
	return s
}

func (s *BotDTO) GetSellPercent() float64 {
	return s.SellPercent
}

func (s *BotDTO) SetUUID(uuid string) *BotDTO {
	s.UUID = uuid
	return s
}

func (s *BotDTO) GetUUID() string {
	return s.UUID
}

func (s *BotDTO) SetCreatedAt(createdAt time.Time) *BotDTO {
	s.CreatedAt = createdAt
	return s
}

func (s *BotDTO) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s *BotDTO) SetUpdatedAt(updatedAt time.Time) *BotDTO {
	s.UpdatedAt = updatedAt
	return s
}

func (s *BotDTO) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

func (s *BotDTO) SetDeletedAt(deletedAt time.Time) *BotDTO {
	s.DeletedAt.Time.Time = deletedAt
	s.DeletedAt.Time.Valid = true
	return s
}

func (s *BotDTO) GetDeletedAt() time.Time {
	return s.DeletedAt.Time.Time
}

func (s *BotDTO) SetActive(active bool) *BotDTO {
	s.Active.Bool = active
	s.Active.Valid = true

	return s
}

func (s *BotDTO) GetActive() bool {
	return s.Active.Bool
}

func (s *BotDTO) SetLimitOrder(limitOrder bool) *BotDTO {
	s.LimitOrder.Bool = limitOrder
	s.LimitOrder.Valid = true

	return s
}

func (s *BotDTO) GetLimitOrder() bool {
	return s.LimitOrder.Bool
}

func (s *BotDTO) SetAutoSell(autoSell bool) *BotDTO {
	s.AutoSell.Bool = autoSell
	s.AutoSell.Valid = true

	return s
}

func (s *BotDTO) GetAutoSell() bool {
	return s.AutoSell.Bool
}

func (s *BotDTO) SetLimitSellPercent(limitSellPercent float64) *BotDTO {
	s.LimitSellPercent.Float64 = limitSellPercent
	s.LimitSellPercent.Valid = true

	return s
}

func (s *BotDTO) GetLimitSellPercent() float64 {
	return s.LimitSellPercent.Float64
}

func (s *BotDTO) SetAutoLimitSellPercent(autoLimitSellPercent float64) *BotDTO {
	s.AutoLimitSellPercent.Float64 = autoLimitSellPercent
	s.AutoLimitSellPercent.Valid = true

	return s
}

func (s *BotDTO) GetAutoLimitSellPercent() float64 {
	return s.AutoLimitSellPercent.Float64
}

func (s *BotDTO) SetLimitBuyPercent(limitBuyPercent float64) *BotDTO {
	s.LimitBuyPercent.Float64 = limitBuyPercent
	s.LimitBuyPercent.Valid = true

	return s
}

func (s *BotDTO) GetLimitBuyPercent() float64 {
	return s.LimitBuyPercent.Float64
}

func (s *BotDTO) SetCommissionPercent(percent float64) *BotDTO {
	s.CommissionPercent = percent

	return s
}

func (s *BotDTO) GetCommissionPercent() float64 {
	return s.CommissionPercent
}

func (s *BotDTO) SetOrderCountLimit(orderCountLimit bool) *BotDTO {
	s.OrderCountLimit.Bool = orderCountLimit
	s.OrderCountLimit.Valid = true

	return s
}

func (s *BotDTO) GetOrderCountLimit() bool {
	return s.OrderCountLimit.Bool
}

func (s *BotDTO) SetOrderCount(orderCount int) *BotDTO {
	s.OrderCount = orderCount
	return s
}

func (s *BotDTO) GetOrderCount() int {
	return s.OrderCount
}

func (s *BotDTO) SetExchangeUserKeyID(exchangeUserKeyID int) *BotDTO {
	s.ExchangeUserKeyID = exchangeUserKeyID
	return s
}

func (s *BotDTO) GetExchangeUserKeyID() int {
	return s.ExchangeUserKeyID
}

func (s *BotDTO) SetAssetType(assetType int) *BotDTO {
	s.AssetType = assetType
	return s
}

func (s *BotDTO) GetAssetType() int {
	return s.AssetType
}
