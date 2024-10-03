package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/ptflp/godecoder"
	"net/http"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/strategy/service"
)

type Strateger interface {
	StrategyCreate(w http.ResponseWriter, r *http.Request)
	StrategyUpdate(w http.ResponseWriter, r *http.Request)
	StrategyGetByID(w http.ResponseWriter, r *http.Request)
	StrategyGetByName(w http.ResponseWriter, r *http.Request)
	StrategyGetList(w http.ResponseWriter, r *http.Request)
	StrategyDelete(w http.ResponseWriter, r *http.Request)
}

type StrategyController struct {
	service service.Strateger
	responder.Responder
	godecoder.Decoder
}

func NewStrategyController(service service.Strateger, components *component.Components) *StrategyController {
	return &StrategyController{service: service, Responder: components.Responder, Decoder: components.Decoder}
}

func (s *StrategyController) StrategyCreate(w http.ResponseWriter, r *http.Request) {
	var strategy StrategyCreateRequest
	err := s.Decode(r.Body, &strategy)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	out := s.service.Create(r.Context(), service.StrategyCreateIn{
		ID:          strategy.ID,
		Name:        strategy.Name,
		UUID:        strategy.UUID,
		Description: strategy.Description,
		ExchangeID:  strategy.ExchangeID,
		Bots:        strategy.Bots,
	})

	if out.ErrorCode != errors.NoError {
		s.OutputJSON(w, StrategyCreateResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	s.OutputJSON(w, StrategyCreateResponse{
		StrategyID: out.StrategyID,
		ErrorCode:  out.ErrorCode,
	})
}

func (s *StrategyController) StrategyUpdate(w http.ResponseWriter, r *http.Request) {
	var req StrategyUpdateRequest
	var res StrategyDefaultResponse

	err := s.Decode(r.Body, &req)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}
	out := s.service.Update(r.Context(), service.StrategyUpdateIn{
		Strategy: req.Strategy,
		Fields:   req.Fields,
	})
	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		s.OutputJSON(w, res)
		return
	}

	res.Success = true
	s.OutputJSON(w, res)
}

func (s *StrategyController) StrategyGetByID(w http.ResponseWriter, r *http.Request) {
	var res StrategyResponse

	idRaw := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	out := s.service.GetByID(r.Context(), service.StrategyGetByIDIn{ID: id})

	if out.ErrorCode != errors.NoError {
		s.OutputJSON(w, StrategyDefaultResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	res.Strategy = out.Strategy
	s.OutputJSON(w, res)

}

func (s *StrategyController) StrategyGetByName(w http.ResponseWriter, r *http.Request) {
	var res StrategyResponse

	name := chi.URLParam(r, "name")

	out := s.service.GetByName(r.Context(), service.StrategyGetByNameIn{Name: name})

	if out.ErrorCode != errors.NoError {
		s.OutputJSON(w, StrategyDefaultResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	res.Strategy = out.Strategy
	s.OutputJSON(w, res)
}

func (s *StrategyController) StrategyGetList(w http.ResponseWriter, r *http.Request) {
	var res StrategiesResponse

	out := s.service.GetList(r.Context())
	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		s.OutputJSON(w, res)
		return
	}

	res.Strategy = out.Strategy
	s.OutputJSON(w, res)
}

func (s *StrategyController) StrategyDelete(w http.ResponseWriter, r *http.Request) {
	var res StrategyDefaultResponse

	idRaw := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	out := s.service.Delete(r.Context(), service.StrategyDeleteIn{ID: id})

	if out.ErrorCode != errors.NoError {
		res.ErrorCode = out.ErrorCode
		s.OutputJSON(w, res)
		return
	}

	res.Success = true
	s.OutputJSON(w, res)
}
