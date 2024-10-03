package controller

import (
	"github.com/ptflp/godecoder"
	"net/http"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
)

type Exchanger interface {
	GetBalanceAccount(w http.ResponseWriter, r *http.Request)
	GetCandles(w http.ResponseWriter, r *http.Request)
}

type Client struct {
	service service.Exchanger
	responder.Responder
	godecoder.Decoder
}

func NewClient(service service.Exchanger, components *component.Components) Exchanger {
	return &Client{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Проверка баланса пользователя
// @Security ApiKeyAuth
// @Tags exchange
// @ID exchangeBalanceAccount
// @Accept  json
// @Produce  json
// @Success 200 {object} GetAccountBalanceOut
// @Router /api/1/exchange/balance [get]
func (p *Client) GetBalanceAccount(w http.ResponseWriter, r *http.Request) {
	out := p.service.GetBalances(r.Context())
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, GetAccountBalanceOut{
			ErrorCode: out.ErrorCode,
		})
		return
	}
	p.OutputJSON(w, GetAccountBalanceOut{
		Success: true,
		DataSpot: BalanceInfo{
			Message:  "SPOT",
			Balances: out.DataSpotBalance,
		},
		DataMargin: BalanceInfo{
			Message:  "Margin",
			Balances: out.DataMarginBalance,
		},
	})
}

// @Summary Получение свечей
// @Security ApiKeyAuth
// @Tags exchange
// @ID exchangeGetCandles
// @Accept  json
// @Produce  json
// @Param symbol query string true "Symbol"
// @Param interval query string true "Interval"
// @Param limit query string true "Limit"
// @Param startTime query string true "StartTime"
// @Param endTime query string true "EndTime"
// @Success 200 {object} GetAccountBalanceOut
// @Router /api/1/exchange/candles [get]
func (p *Client) GetCandles(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	interval := r.URL.Query().Get("interval")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	startTime, err := strconv.ParseInt(r.URL.Query().Get("startTime"), 10, 64)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	endTime, err := strconv.ParseInt(r.URL.Query().Get("endTime"), 10, 64)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.GetCandles(r.Context(), service.GetCandlesIn{
		Symbol:    symbol,
		Interval:  service.KlineInterval(interval),
		Limit:     limit,
		StartTime: startTime,
		EndTime:   endTime,
	})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, CandlesResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}
	p.OutputJSON(w, CandlesResponse{

		Candles: out.Candles,
	})

}
