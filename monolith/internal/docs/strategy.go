package docs

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/strategy/controller"
)

// swagger:route POST /api/1/strategy/create strategy strategyCreateRequest
// Создание стратегии.
// security:
//      - Bearer: []
// responses:
//  200: strategyCreateResponse
//  400: description: Bad request
//  500: description: Internal server error

// swagger:parameters strategyCreateRequest
type strategyCreateRequest struct {
	// in:body
	Body controller.StrategyCreateRequest
}

// swagger:response strategyCreateResponse
type strategyCreateResponse struct {
	//  in:body
	Body controller.StrategyCreateResponse
}

// swagger:route POST /api/1/strategy/update strategy strategyUpdateRequest
// Редактирование стратегии.
// security:
//      - Bearer: []
// responses:
//  200: strategyUpdateResponse
//  400: description: Bad request
//  500: description: Internal server error

// swagger:parameters strategyUpdateRequest
type strategyUpdateRequest struct {
	// in:body
	Body controller.StrategyUpdateRequest
}

// swagger:parameters strategyUpdateResponse
type strategyUpdateResponse struct {
	// in:body
	Body controller.StrategyDefaultResponse
}

// swagger:route DELETE /api/1/strategy/{id} strategy strategyDeleteRequest
// Удаление стратегии.
// security:
//      - Bearer: []
// responses:
//  200: strategyDeleteResponse
//  400: description: Bad request
//  500: description: Internal server error

// swagger:parameters strategyDeleteRequest
type strategyDeleteRequest struct {
	// strategy_id
	//
	// required:true
	// in:path
	ID string `json:"id"`
}

// swagger:response strategyDeleteResponse
type strategyDeleteResponse struct {
	// in:body
	Body controller.StrategyDefaultResponse
}

// swagger:route GET /api/1/strategy/{id} strategy strategyGetByIDRequest
// Получение стратегии по айди.
// security:
//      - Bearer: []
// responses:
//  200: strategyGetByIDResponse
//  400: description: Bad request
//  500: description: Internal server error

// swagger:parameters strategyGetByIDRequest
type strategyGetByIDRequest struct {
	// strategy_id
	//
	// required:true
	// in:path
	ID string `json:"id"`
}

// swagger:response strategyGetByIDResponse
type strategyGetByIDResponse struct {
	// in:body
	Body controller.StrategyResponse
}

// swagger:route GET /api/1/strategy/{name} strategy strategyGetByNameRequest
// Получение стратегии по имени.
// security:
//      - Bearer: []
// responses:
//  200: strategyGetByNameResponse
//  400: description: Bad request
//  500: description: Internal server error

// swagger:parameters strategyGetByNameRequest
type strategyGetByNameRequest struct {
	// strategy_name
	//
	// required:true
	// in:path
	Name string `json:"name"`
}

// swagger:response strategyGetByNameResponse
type strategyGetByNameResponse struct {
	// in:body
	Body controller.StrategyResponse
}

// swagger:route GET /api/1/strategy/list strategy strategyGetListRequest
// Получение списка стратегий.
// security:
//      - Bearer: []
// responses:
//  200: strategyGetListResponse
//  400: description: Bad request
//  500: description: Internal server error

// swagger:response strategyGetListResponse
type strategyGetListResponse struct {
	// in:body
	Body controller.StrategiesResponse
}
