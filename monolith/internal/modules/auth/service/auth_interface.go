package service

import "context"

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Auther
type Auther interface {
	Register(ctx context.Context, in RegisterIn, field int) RegisterOut
	AuthorizeEmail(ctx context.Context, in AuthorizeEmailIn) AuthorizeOut
	AuthorizeRefresh(ctx context.Context, in AuthorizeRefreshIn) AuthorizeOut
	AuthorizePhone(ctx context.Context, in AuthorizePhoneIn) AuthorizeOut
	SendPhoneCode(ctx context.Context, in SendPhoneCodeIn) SendPhoneCodeOut
	VerifyEmail(ctx context.Context, in VerifyEmailIn) VerifyEmailOut
	SocialCallback(ctx context.Context, in SocialCallbackIn) AuthorizeOut
	SocialGetRedirectURL(ctx context.Context, in SocialGetRedirectUrlIn) SocialGetRedirectUrlOut
}

const (
	RegisterEmail = iota + 1
	RegisterPhone
)

type VerifyEmailIn struct {
	Hash  string
	Email string
}

type VerifyEmailOut struct {
	Success   bool
	ErrorCode int
}

type SendPhoneCodeIn struct {
	Phone string
}

type SendPhoneCodeOut struct {
	Phone     string
	Code      int
	ErrorCode int
}

type AuthorizeIn struct {
	Email    string
	Password string
}

type AuthorizeOut struct {
	UserID       int
	AccessToken  string
	RefreshToken string
	ErrorCode    int
}

type RegisterIn struct {
	Email          string
	Phone          string
	Password       string
	IdempotencyKey string
}

type RegisterOut struct {
	Status    int
	ErrorCode int
}

type AuthorizeEmailIn struct {
	Email          string
	Password       string
	RetypePassword string
}

type AuthorizeRefreshIn struct {
	UserID int
}

type AuthorizePhoneIn struct {
	Phone string
	Code  int
}

type Out struct {
	ErrorCode int
}

type SocialCallbackIn struct {
	Code     string
	Provider string
}

type SocialGetRedirectUrlIn struct {
	Provider string
}
type SocialGetRedirectUrlOut struct {
	Url       string
	ErrorCode int
}
