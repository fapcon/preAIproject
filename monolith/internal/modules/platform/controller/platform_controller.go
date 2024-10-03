package controller

import (
	"github.com/ptflp/godecoder"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/handlers"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/platform/service"
)

type Platformer interface {
	BotInfo(w http.ResponseWriter, r *http.Request)
}

type Platform struct {
	platform *service.Platform
	responder.Responder
	godecoder.Decoder
}

func NewPlatform(platform *service.Platform, components *component.Components) Platformer {
	return &Platform{platform: platform, Responder: components.Responder, Decoder: components.Decoder}
}

func (p *Platform) BotInfo(w http.ResponseWriter, r *http.Request) {
	userClaims, err := handlers.ExtractUser(r)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	var req BotInfoRequest

	err = p.Decode(r.Body, &req)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.platform.GetBotInfo(r.Context(), service.GetBotInfoIn{
		BotUUID: req.BotUUID,
		UserID:  userClaims.ID,
	})
	if out.ErrorCode != errors.NoError {
		p.OutputJSON(w, BotInfoResponse{
			ErrorCode: out.ErrorCode,
		})
		return
	}

	p.OutputJSON(w, BotInfoResponse{
		Success: true,
		Data: BotInfoData{
			Bot:      out.Data.Bot,
			Orders:   out.Data.Orders,
			Webhooks: out.Data.Webhooks,
		},
	})
}
