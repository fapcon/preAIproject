//go:build !test

package docs

import (
	ucontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/controller"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/service"
)

//go:generate swagger generate spec -o ../../static/swagger.json --scan-models

// swagger:route GET /api/1/user/profile user profileRequest
// Получение профиля пользователя.
// security:
//   - Bearer: []
// responses:
//   200: profileResponse

// swagger:response profileResponse
type profileResponse struct {
	// in:body
	Body ucontroller.ProfileResponse
}

// swagger:route POST /api/1/user/changePassword user changePasswordRequest
// Изменение пароля пользователя.
// security:
//   - Bearer: []
// responses:
//   200: changePasswordResponse

// swagger:parameters changePasswordRequest
type ChangePasswordRequest struct {
	// in:body
	Body service.ChangePasswordIn
}

// swagger:response changePasswordResponse
type ChangePasswordResponse struct {
	// Password changed successfully!
	// in:body
	Body service.ChangePasswordOut
}

// swagger:route POST /api/1/user/resetPassword user resetPasswordRequest
// Сброс пароля пользователя.
// security:
//   - Bearer: []
// responses:
//   200: resetPasswordResponse

// swagger:parameters resetPasswordRequest
type ResetPasswordRequest struct {
	// in:body
	Body service.ResetPasswordIn
}

// swagger:response resetPasswordResponse
type ResetPasswordResponse struct {
	// Password reset successfully!
	// in:body
	Body service.ResetPasswordOut
}

// swagger:route POST /api/1/user/sendCode user sendCodeRequest
// Отправка кода сброса по электронной почте.
// security:
//   - Bearer: []
// responses:
//   200: sendCodeResponse

// swagger:parameters sendCodeRequest
type SendCodeRequest struct {
	// in:body
	Body service.SendResetCodeEmailIn
}

// swagger:response sendCodeResponse
type SendCodeResponse struct {
	// Reset code email sent successfully!
	// in:body
	Body service.SendResetCodeEmailOut
}
