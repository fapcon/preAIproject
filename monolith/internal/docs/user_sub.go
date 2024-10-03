package docs

import (
	uscontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user_sub/controller"
)

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route POST /api/1/user/subscription/add user userSubAddRequest
// Добавление подписки на пользователя.
// security:
//   - Bearer: []
// responses:
//   200: userSubAddResponse

//swagger:parameters userSubAddRequest
type subAddRequest struct {
	//in: body
	Body uscontroller.UserSubAddRequest
}

//swagger:response userSubAddResponse
type userSubAddResponse struct {
	// in: body
	Body uscontroller.UserSubResponse
}

// swagger:route GET /api/1/user/subscription/list user userSubListRequest
// Список подписок пользователя.
// security:
//   - Bearer: []
// responses:
//   200: userSubListResponse

//swagger:response userSubListResponse
type userSubListResponse struct {
	// in: body
	Body uscontroller.UserSubListResponse
}

// swagger:route DELETE /api/1/user/subscription/delete user userSubDeleteRequest
// Удаление подписки на пользователя.
// security:
//   - Bearer: []
// responses:
//   200: userSubDeleteResponse

//swagger:parameters userSubDeleteRequest
type userSubDeleteRequest struct {
	//in: body
	Body uscontroller.UserSubDelRequest
}

//swagger:response userSubDeleteResponse
type userSubDeleteResponse struct {
	// in: body
	Body uscontroller.UserSubResponse
}
