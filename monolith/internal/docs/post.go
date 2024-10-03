package docs

import "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/posts/controller"

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route POST /api/1/posts/create posts postCreateRequest
// Создание поста.
// security:
//   - Bearer: []
// responses:
//   200: postCreateResponse

// swagger:parameters postCreateRequest
type postCreateRequest struct {
	// in:body
	Body controller.PostUserCreateRequest
}

// swagger:response postCreateResponse
type postCreateResponse struct {
	// in:body
	Body controller.PostCreateResponse
}

// swagger:route POST /api/1/posts/delete posts postDeleteRequest
// Удалить пост.
// security:
//   - Bearer: []
// responses:
//   200: postDeleteResponse

// swagger:parameters postDeleteRequest
type postDeleteRequest struct {
	// in:body
	Body controller.PostUserDeletedRequest
}

// swagger:response postDeleteResponse
type postDeleteResponse struct {
	// in:body
	Body controller.PostDeleteResponse
}

// swagger:route POST /api/1/posts/update posts postUpdateRequest
// Обновить пост.
// security:
//   - Bearer: []
// responses:
//   200: postUpdateResponse

// swagger:parameters postUpdateRequest
type postUpdateRequest struct {
	// in:body
	Body controller.PostUserUpdateRequest
}

// swagger:response postUpdateResponse
type postUpdateResponse struct {
	// in:body
	Body controller.PostUpdateResponse
}

// swagger:route GET /api/1/posts/id posts postGetByIdRequest
// Поиск поста по id.
// security:
//   - Bearer: []
// responses:
//   200: postGetByIdResponse

// swagger:parameters postGetByIdRequest
type postGetByIdRequest struct {
	// in:query
	Id string `json:"id"`
}

// swagger:response postGetByIdResponse
type postGetByIdResponse struct {
	// in:body
	Body controller.PostUpdateResponse
}

// swagger:route GET /api/1/posts/tape posts postGetAllPostRequest
// Вывод постов в ленте.
// security:
// - Bearer: []
// responses:
// 	200: postGetAllResponse

// swagger:parameters postGetAllPostRequest
type postGetAllPostRequest struct {
	// in:query
	Limit int `json:"limit"`
	// in:query
	Offset int `json:"offset"`
}

// swagger:response postGetAllPostResponse
type postGetAllPostResponse struct {
	// in:body
	Body controller.PostListTapeResponse
}
