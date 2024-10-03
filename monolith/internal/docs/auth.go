package docs

import (
	acontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/controller"
)

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route POST /api/1/auth/register auth registerRequest
// Регистрация пользователя.
// responses:
//   200: registerResponse

// swagger:parameters registerRequest
type registerRequest struct {
	// in:body
	Body acontroller.RegisterRequest
}

// swagger:response registerResponse
type registerResponse struct {
	// in:body
	Body acontroller.RegisterResponse
}

// swagger:route POST /api/1/auth/login auth loginRequest
// Авторизация пользователя.
// responses:
//   200: loginResponse

// swagger:parameters loginRequest
type loginRequest struct {
	// in:body
	Body acontroller.LoginRequest
}

// swagger:response loginResponse
type loginResponse struct {
	// in:body
	Body acontroller.AuthResponse
}

// swagger:route POST /api/1/auth/refresh auth refreshRequest
// Обновление рефреш токена.
// security:
//   - Bearer: []
// responses:
//   200: refreshResponse

// swagger:response refreshResponse
type refreshRespone struct {
	// in:body
	Body acontroller.AuthResponse
}

// swagger:route POST /api/1/auth/verify auth verifyRequest
// Верификация почты/телефона пользователя.
// responses:
//   200: verifyResponse

// swagger:parameters verifyRequest
type verifyRequest struct {
	// in:body
	Body acontroller.VerifyRequest
}

// swagger:response verifyResponse
type verifyResponse struct {
	// in:body
	Body acontroller.AuthResponse
}
