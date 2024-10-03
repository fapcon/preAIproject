package tg_bot

import tg_bot "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

type TgBot interface {
	SendMessage(text string) error
	Run(errChan <-chan tg_bot.ErrMessage)
	RunNotification(NotificationChan <-chan tg_bot.NotificationMessage)
}
