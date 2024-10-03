package controller

import "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/service"

type EMA_periodsResponse struct {
	ErrorCode int
	EMA       service.EMA_out
}

type EMA_periodsRequest struct {
	FirstEMA  int `json:"firstEMA"`
	SecondEMA int `json:"secondEMA"`
	ThirdEMA  int `json:"thirdEMA"`
}

type DynamicPairsResponse struct {
	ErrorCode int
	Symbols   []string
}
