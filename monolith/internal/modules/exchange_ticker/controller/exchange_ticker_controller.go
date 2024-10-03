package controller

import (
	"net/http"

	"github.com/ptflp/godecoder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/service"
)

type ExchangeTicker interface {
	Ticker(http.ResponseWriter, *http.Request)
}

type Ticker struct {
	ticker service.ExchangeTicker
	responder.Responder
	godecoder.Decoder
}

func NewTicker(ticker service.ExchangeTicker, components *component.Components) ExchangeTicker {
	return &Ticker{ticker: ticker, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Получение курса криптовалют
// @Security ApiKeyAuth
// @Tags exchange
// @ID ticker
// @Accept  json
// @Produce  json
// @Success 200 {object} TickerResponse
// @Router /api/1/exchange/ticker [get]
func (p *Ticker) Ticker(w http.ResponseWriter, r *http.Request) {
	var res TickerResponse

	out := p.ticker.GetTicker(r.Context())

	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		p.OutputJSON(w, res)
		return
	}

	p.OutputJSON(w, TickerResponse{
		Success: true,
		Data:    out.Data,
	})
}
