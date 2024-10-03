package provider

import (
	"fmt"
	"go.uber.org/zap"
	"net/smtp"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
)

type Sender interface {
	Send(in SendIn) error
}

type SendIn struct {
	To    string
	From  string
	Title string
	Type  int
	Data  []byte
}

const (
	textPlain = "text/plain"
	textHtml  = "text/html"

	TextPlain = iota + 1
	TextHtml
)

type Email struct {
	conf   config.Email
	client smtp.Auth
	addr   string
	logger *zap.Logger
}

func NewEmail(conf config.Email, logger *zap.Logger) *Email {
	emailAuth := smtp.PlainAuth("", conf.Credentials.Login, conf.Credentials.Password, conf.Credentials.Host)

	return &Email{conf: conf, client: emailAuth, addr: fmt.Sprintf("%s:%s", conf.Credentials.Host, conf.Port), logger: logger}
}

func (e *Email) Send(in SendIn) error {
	emailBody := string(in.Data)
	var contentType string

	switch in.Type {
	case TextPlain:
		contentType = textPlain
	case TextHtml:
		contentType = textHtml
	default:
		contentType = textPlain
	}

	mime := "MIME-version: 1.0;\nContent-Type: " + contentType + "; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + in.Title + "\n"
	msg := []byte(subject + mime + "\n" + emailBody)

	if err := smtp.SendMail(e.addr, e.client, e.conf.From, []string{in.To}, msg); err != nil {
		e.logger.Error("email: sent msg err", zap.Error(err))
		return err
	}

	return nil
}
