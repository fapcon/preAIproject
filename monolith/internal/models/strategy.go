package models

//go:generate easytags $GOFILE json,mapper
type Strategy struct {
	ID          int    `json:"id" mapper:"id"`
	Name        string `json:"name" mapper:"name"`
	UUID        string `json:"uuid" mapper:"uuid"`
	Description string `json:"description" mapper:"description"`
	ExchangeID  int    `json:"exchange_id" mapper:"exchange_id"`
	//UserID      int    `json:"user_id" mapper:"user_id"`
	//APIKey      string `json:"api_key" mapper:"api_key"`  TODO: нужно ли?
	//SecretKey   string `json:"secret_key" mapper:"secret_key"`
	Bots []Bot `json:"bots" mapper:"bots"`
}
