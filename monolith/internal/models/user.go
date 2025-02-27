package models

type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Password      string `json:"-"`
	Role          int    `json:"role"`
	Status        int    `json:"status"`
	Verified      bool   `json:"verified"`
	EmailVerified bool   `json:"email_verified"`
	PhoneVerified bool   `json:"phone_verified"`
}
