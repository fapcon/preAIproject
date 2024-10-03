package docs

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/platform/controller"
)

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models
// swagger:route POST /platform/hook platform platformHookRequest
// Обработка вебхуков пользователя.
// responses:
//   200: platformHookResponse

// swagger:parameters platformHookRequest
type platformHookRequest struct {
	// in:query
	Body string
}

// swagger:response platformHookResponse
type platformHookResponse struct {
	// in:body
	Body controller.WebhookProcessResponse
}
