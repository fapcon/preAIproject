package storages

import (
	"gitlab.com/golight/orm/db/adapter"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/cache"
	vstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/storage"
	sstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/storage"
	comment_storage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/comments/storage"
	istorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/storage"
	lstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/storage"
	ostorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/storage"
	tstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/storage"
	kstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/storage"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/platform/storage"
	pstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/posts/storage"
	statstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/statistics/storage"
	ststorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/strategy/storage"
	ustorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/storage"
	usubstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user_sub/storage"
	wstorage "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/webhook_process/storage"
)

type Storages struct {
	User                  ustorage.Userer
	Bot                   sstorage.Boter
	Ticker                tstorage.ExchangeTicker
	ExchangeList          lstorage.ExchangeLister
	ExchangeUserKey       kstorage.ExchangeUserKeyer
	ExchangeOrder         ostorage.ExchangeOrderer
	ExchangeOrderLog      ostorage.ExchangeOrderLogger
	StrategyPair          sstorage.StrategyPairer
	WebhookProcess        wstorage.WebhookProcesser
	WebhookProcessHistory wstorage.WebhookProcessHistorer
	Strategy              ststorage.Strateger
	StrategySubscribers   storage.StrategySubscriberer
	Indicator             istorage.Indicatorer
	Post                  pstorage.Poster
	Comment               comment_storage.Commenter
	UserSub               usubstorage.UserSuber
	Verifier              vstorage.Verifier
	Statistics            statstorage.Statisticer
}

func NewStorages(sqlAdapter *adapter.SQLAdapter, cache cache.Cache) *Storages {
	return &Storages{
		User:                  ustorage.NewUserStorage(sqlAdapter, cache),
		Bot:                   sstorage.NewBotStorage(*sqlAdapter, cache),
		Ticker:                tstorage.NewTicker(*sqlAdapter),
		ExchangeList:          lstorage.NewExchangeList(*sqlAdapter),
		ExchangeUserKey:       kstorage.NewExchangeUserKey(*sqlAdapter, cache),
		ExchangeOrder:         ostorage.NewExchangeOrder(*sqlAdapter),
		ExchangeOrderLog:      ostorage.NewExchangeOrderLog(*sqlAdapter),
		StrategyPair:          sstorage.NewStrategyPairStorage(*sqlAdapter),
		WebhookProcess:        wstorage.NewWebhookProcess(*sqlAdapter),
		WebhookProcessHistory: wstorage.NewWebhookProcessHistory(*sqlAdapter),
		Strategy:              ststorage.NewStrategy(sqlAdapter),
		StrategySubscribers:   storage.NewStrategySubscribers(*sqlAdapter),
		Indicator:             istorage.NewIndicatorStorage(),
		Comment:               comment_storage.NewComment(sqlAdapter),
		Post:                  pstorage.NewPostStorage(sqlAdapter, cache),
		UserSub:               usubstorage.NewUserSubStorage(sqlAdapter),
		Verifier:              vstorage.NewEmailVerify(sqlAdapter),
		Statistics:            statstorage.NewStatisticStorage(*sqlAdapter, cache),
	}
}
