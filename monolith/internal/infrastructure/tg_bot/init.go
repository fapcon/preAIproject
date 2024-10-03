package tg_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"strconv"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
)

type TgBotClient struct {
	bot       *tgbotapi.BotAPI
	logger    *zap.Logger
	trustedID int64
}

func NewTgBotClient(config config.AppConf, logger *zap.Logger) TgBot {
	bot, err := tgbotapi.NewBotAPI(config.TgBot.Token)
	if err != nil {
		logger.Error("error while initialisation TgBotClient", zap.Error(err))
	}

	chatID, err := strconv.Atoi(config.TgBot.TrustedChatID)
	if err != nil {
		logger.Error("TrustedChatID should be convertable to int", zap.Error(err))
	}

	return &TgBotClient{
		bot:       bot,
		logger:    logger,
		trustedID: int64(chatID),
	}
}

func (t *TgBotClient) SendMessage(text string) error {
	msg := tgbotapi.NewMessage(t.trustedID, text)
	_, err := t.bot.Send(msg)
	return err
}
