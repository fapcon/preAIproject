package docks

//swagger:route Post /api/v1/user/register auth registerRequest
// Регистрация пользователя.
// Responses:
//   200: registerResponse

//swagger:parameters registerRequest
type registerRequest struct {
	// Userdata - данные пользователя
	// in: body
	// required: true
	// example: {"email":"solo228@gmail.com","password":"322"}
	Userdata string
}

// swagger:response registerResponse
type registerResponse struct {
	// in: body
	// UserID айди пользователя.
	UserID int
}

//swagger:route Post /api/v1/user/login auth loginRequest
// Авторизация пользователя.
// Responses:
//   200: loginResponse

//swagger:parameters loginRequest
type loginRequest struct {
	// Userdata - данные пользователя
	// in: body
	// required: true
	// example: {"email":"solo228@gmail.com","password":"322"}
	Userdata string
}

// swagger:response loginResponse
type loginResponse struct {
	// in: body
	// User профиль пользователя.
	User string
}

//swagger:route Get /api/v1/user/profile auth profileRequest
// Профиль пользователя.
// Responses:
//   200: profileResponse

//swagger:parameters profileRequest

// swagger:response profileResponse
type profileResponse struct {
	// in: body
	// User профиль пользователя.
	User string
}

//swagger:route Get /api/v1/user/logout auth logoutRequest
// Выйти из акаунта.
// Responses:
//   200: logoutResponse

//swagger:parameters logoutRequest

//swagger:response logoutResponse
type logoutResponse struct {
	// in: body
	// Nothing заглушка.
	Nothing string
}

//swagger:route Get /api/v1/user/{provider}/login auth SocialGetRedirectURLRequest
// Получить URL для авторизации.
// Responses:
//   200: SocialGetRedirectURLResponse

//swagger:parameters SocialGetRedirectURLRequest
type SocialGetRedirectURLRequest struct {
	// Provider - способ авторизации для данного пользователя
	// in: path
	// required: true
	// example: google, facebook
	Provider string `json:"provider"`
}

// swagger:response SocialGetRedirectURLResponse
type SocialGetRedirectURLResponse struct {
	// in: body
	// Url ссылка для редиректа.
	Url string
}
