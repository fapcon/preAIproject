package controller

import (
	"github.com/ptflp/godecoder"
	"net/http"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/service"
)

type Indicatorer interface {
	EMA(w http.ResponseWriter, r *http.Request)
	GetDynamicPairs(w http.ResponseWriter, r *http.Request)
}

type Indicator struct {
	service service.Indicatorer
	responder.Responder
	godecoder.Decoder
}

func NewIndicator(service service.Indicatorer, components *component.Components) Indicatorer {
	return &Indicator{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Получение индикатора EMA
// @Security ApiKeyAuth
// @Tags exchange
// @ID exchangeEMA
// @Accept  json
// @Produce  json
// @Param symbol query string true "Symbol (ETHBTC,LTCBTC,BNBBTC,NEOBTC,QTUMETH,EOSETH,SNTETH,BNTETH ...)"
// @Param interval query string true "Interval (1m,3m,5m,15m,30m,1h,2h,4h,6h,8h,12h,1d,3d,1w,1M)"
// @Param limit query string true "Limit"
// @Param object body EMA_periodsRequest true "EMA_periodsRequest"
// @Success 200 {object} EMA_periodsResponse
// @Router /api/1/exchange/ema [post]
func (i *Indicator) EMA(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	interval := r.URL.Query().Get("interval")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	var req EMA_periodsRequest
	err := i.Decode(r.Body, &req)
	if err != nil {
		i.ErrorBadRequest(w, err)
		return
	}
	out := i.service.EMA(symbol, interval, limit, req.FirstEMA, req.SecondEMA, req.ThirdEMA, r.Context())
	if out.ErrorCode != errors.NoError {
		i.OutputJSON(w, EMA_periodsResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}
	i.OutputJSON(w, EMA_periodsResponse{
		EMA: out,
	})
}

// @Summary Получение торгующих пар с биржи бинанс
// @Security ApiKeyAuth
// @Tags exchange
// @ID exchangePairs
// @Accept  json
// @Produce  json
// @Success 200 {object} DynamicPairsResponse
// @Router /api/1/exchange/pairs [get]
func (i *Indicator) GetDynamicPairs(w http.ResponseWriter, r *http.Request) {
	out := i.service.GetDynamicPairBinance(r.Context())
	if out.ErrorCode != errors.NoError {
		i.OutputJSON(w, DynamicPairsResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}
	i.OutputJSON(w, DynamicPairsResponse{
		Symbols: out.Symbols,
	})
}
