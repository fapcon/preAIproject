package docs

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/statistics/controller"
)

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route GET /api/1/statistics/user statistics userStatisticRequest
// Получение статистики всех ботов пользователя.
// security:
//   - Bearer: []
// responses:
//   200: userStatisticResponse

// swagger:response userStatisticResponse
type userStatisticResponse struct {
	// in:body
	Body controller.UserStatisticResponse
}

// swagger:route POST /api/1/statistics/bot statistics botStatisticRequest
// Получение статистики бота по идентификатору.
// responses:
//   200: botStatisticResponse

// swagger:parameters botStatisticRequest
type botStatisticRequest struct {
	// in:body
	Body controller.BotStatisticRequest
}

// swagger:response botStatisticResponse
type botStatisticResponse struct {
	// in:body
	Body controller.BotStatisticResponse
}

// swagger:route POST /api/1/statistics/delete statistics botStatisticDeleteRequest
// Удаление статистики бота по идентификатору.
// responses:
//   200: botStatisticDeleteResponse

// swagger:parameters botStatisticDeleteRequest
type botStatisticDeleteRequest struct {
	// in:body
	Body controller.BotStatisticDeleteRequest
}

// swagger:response botStatisticDeleteResponse
type botStatisticDeleteResponse struct {
	// in:body
	Body controller.BotStatisticDeleteResponse
}
