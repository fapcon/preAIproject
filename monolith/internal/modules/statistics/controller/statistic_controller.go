package controller

import (
	"github.com/ptflp/godecoder"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/statistics/service"
)

type Statisticer interface {
	GetBotStatistic(w http.ResponseWriter, r *http.Request)
	GetUserStatistic(w http.ResponseWriter, r *http.Request)
	DeleteStatistic(w http.ResponseWriter, r *http.Request)
}

type Statistic struct {
	service service.Statisticer
	responder.Responder
	godecoder.Decoder
}

func NewStatisticController(service service.Statisticer, components *component.Components) Statisticer {
	return &Statistic{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

// @Summary Получение статистики бота по идентификатору.
// @Security ApiKeyAuth
// @Tags statistics
// @ID botStatisticsGet
// @Accept  json
// @Produce  json
// @Param object body BotStatisticRequest true "BotStatisticRequest"
// @Success 200 {object} BotStatisticResponse
// @Router /api/1/statistics/bot [post]
func (s *Statistic) GetBotStatistic(w http.ResponseWriter, r *http.Request) {
	var req BotStatisticRequest
	var res BotStatisticResponse
	err := s.Decode(r.Body, &req)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	out := s.service.GetBotStatistic(r.Context(), service.StatisticIn{
		BotUUID: req.BotUUID,
	})

	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		s.OutputJSON(w, res)
		return
	}

	s.OutputJSON(w, BotStatisticResponse{
		Success: true,
		Data:    out.Statistic,
	})
}

// @Summary Получение статистики всех ботов пользователя.
// @Security ApiKeyAuth
// @Tags statistics
// @ID userStatisticsGet
// @Accept  json
// @Produce  json
// @Success 200 {object} UserStatisticResponse
// @Router /api/1/statistics/user [get]
func (s *Statistic) GetUserStatistic(w http.ResponseWriter, r *http.Request) {
	var res UserStatisticResponse

	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	out := s.service.GetUserBotStatistic(r.Context(), service.UserStatisticIn{
		UserID: userClaims.ID,
	})

	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		s.OutputJSON(w, res)
		return
	}

	res.Success = true
	res.Data = out.Statistic
	s.OutputJSON(w, res)

}

// @Summary Удаление статистики бота по идентификатору.
// @Security ApiKeyAuth
// @Tags statistics
// @ID botStatisticsDelete
// @Accept  json
// @Produce  json
// @Param object body BotStatisticDeleteRequest true "BotStatisticDeleteRequest"
// @Success 200 {object} BotStatisticDeleteResponse
// @Router /api/1/statistics/delete [post]
func (s *Statistic) DeleteStatistic(w http.ResponseWriter, r *http.Request) {
	var req BotStatisticDeleteRequest
	var res BotStatisticDeleteResponse

	err := s.Decode(r.Body, &req)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	out := s.service.DeleteBotStatistic(r.Context(), service.StatisticIn{
		BotUUID: req.BotUUID,
	})

	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		s.OutputJSON(w, res)
		return
	}

	res.Success = true
	s.OutputJSON(w, res)
}
