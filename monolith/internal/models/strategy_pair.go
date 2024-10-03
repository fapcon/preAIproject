package models

type StrategyPair struct {
	ID         int `json:"id" db:"id" db_type:"BIGSERIAL primary key" db_default:"not null" mapper:"id"`
	StrategyID int `json:"strategy_id" db:"strategy_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"strategy_id"`
	PairID     int `json:"pair_id" db:"pair_id" db_ops:"create,update" db_type:"int" db_default:"not null" mapper:"pair_id"`
}
