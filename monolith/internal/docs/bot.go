package docs

import (
	scontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/controller"
)

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route GET /api/1/bot/create bot bot
// Создание бота.
// security:
//   - Bearer: []
// responses:
//   200: botCreateResponse

// swagger:response botCreateResponse
type botCreateResponse struct {
	// in:body
	Body scontroller.BotResponse
}

// swagger:route POST /api/1/bot/update bot botUpdateRequest
// Редактирование бота.
// responses:
//   200: botUpdateResponse

// swagger:parameters botUpdateRequest
type botUpdateRequest struct {
	// in:body
	Body scontroller.UpdateRequest
}

// swagger:response botUpdateResponse
type botUpdateResponse struct {
	// in:body
	Body scontroller.UpdateResponse
}

// swagger:route POST /api/1/bot/delete bot botDeleteRequest
// Удаление бота.
// responses:
//   200: botDeleteResponse

// swagger:parameters botDeleteRequest
type botDeleteRequest struct {
	// in:body
	Body scontroller.BotUUIDRequest
}

// swagger:response botDeleteResponse
type botDeleteResponse struct {
	// in:body
	Body scontroller.DefaultResponse
}

// swagger:route POST /api/1/bot/toggle bot botToggleRequest
// Активация, деактивация бота.
// responses:
//   200: botToggleResponse

// swagger:parameters botToggleRequest
type botToggleRequest struct {
	// in:body
	Body scontroller.BotToggleRequest
}

// swagger:response botToggleResponse
type botToggleResponse struct {
	// in:body
	Body scontroller.DefaultResponse
}

// swagger:route POST /api/1/bot/get bot botGetRequest
// Получение бота по идентификатору.
// responses:
//   200: botGetResponse

// swagger:parameters botGetRequest
type botGetRequest struct {
	// in:body
	Body scontroller.GetBotRequest
}

// swagger:response botGetResponse
type botGetResponse struct {
	// in:body
	Body scontroller.BotData
}

// swagger:route GET /api/1/bot/list bot botList
// Список ботов пользователя.
// security:
//   - Bearer: []
// responses:
//   200: botListResponse

// swagger:response botListResponse
type botListResponse struct {
	// in:body
	Body scontroller.BotListResponse
}

// swagger:route POST /api/1/bot/signals bot webhookSignalsRequest
// Получение хука и сигналов.
// responses:
//   200: webhookSignalsResponse

// swagger:parameters webhookSignalsRequest
type webhookSignalsRequest struct {
	// in:body
	Body scontroller.WebhookSignalRequest
}

// swagger:response webhookSignalsResponse
type webhookSignalsResponse struct {
	// in:body
	Body scontroller.WebhookSignalResponse
}
