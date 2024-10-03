package tg_bot

import tg_bot "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"

func (t *TgBotClient) Run(errChan <-chan tg_bot.ErrMessage) {
	for {
		message := <-errChan
		err := t.SendMessage(message.String())
		if err != nil {
			t.logger.Info("error while sending error to tg channel")
		}
	}
}

func (t *TgBotClient) RunNotification(NotificationChan <-chan tg_bot.NotificationMessage) {
	for {
		message := <-NotificationChan
		err := t.SendMessage(message.String())
		if err != nil {
			t.logger.Info("error while sending NotificationMessage to tg channel")
		}
	}
}
