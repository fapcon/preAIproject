package service

import (
	"go.uber.org/zap"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/provider"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Notifier
type Notifier interface {
	Push(in PushIn) PushOut
}

type Notify struct {
	conf   config.Email
	email  provider.Sender
	logger *zap.Logger
}

func NewNotify(conf config.Email, email provider.Sender, logger *zap.Logger) Notifier {
	return &Notify{email: email, conf: conf, logger: logger}
}

func (n *Notify) Push(in PushIn) PushOut {
	err := n.email.Send(provider.SendIn{
		To:    in.Identifier,
		From:  n.conf.From,
		Title: in.Title,
		Type:  provider.TextPlain,
		Data:  in.Data,
	})
	if err != nil {
		n.logger.Error("send email err", zap.Error(err))
		return PushOut{
			ErrorCode: errors.NotifyEmailSendErr,
		}
	}

	return PushOut{}
}

const (
	PushEmail = iota + 1
	PushPhone
)

type PushIn struct {
	Identifier string
	Type       int
	Title      string
	Data       []byte
	Options    []interface{}
}

type PushOut struct {
	ErrorCode int
}
