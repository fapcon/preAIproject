package modules

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/grpc_clients"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	aservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/service"
	bservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/service"
	comments_service "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/comments/service"
	emetrics "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/metrics"
	cservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"
	iservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/service"
	lservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/service"
	oservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/service"
	tiservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/service"
	kservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/service"
	pservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/platform/service"
	post_service "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/posts/service"
	statservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/statistics/service"
	stservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/strategy/service"
	uservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/service"
	usubservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user_sub/service"
	wservice "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/webhook_process/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/storages"
)

type Services struct {
	User         uservice.Userer
	Auth         aservice.Auther
	Bot          bservice.Boter
	ExchangeList lservice.ExchangeLister
	Order        oservice.ExchangeOrderer
	Ticker       tiservice.ExchangeTicker
	UserKey      kservice.ExchangeUserKeyer
	Platform     *pservice.Platform
	Webhook      wservice.WebhookProcesser
	Indicator    iservice.Indicatorer
	Client       cservice.Exchanger
	Post         post_service.Poster
	Comments     comments_service.CommentService
	UserSub      usubservice.UserSuber
	Strategy     stservice.Strateger
	Statistics   statservice.Statisticer
}

func NewServices(storages *storages.Storages, components *component.Components) *Services {
	user := uservice.NewUserService(storages.User, components.Logger, components)
	userKey := kservice.NewUserKey(storages.ExchangeUserKey, components)
	ticker := tiservice.NewTicker(storages.Ticker, components)
	statistics := statservice.NewStatistic(storages.Statistics)
	order := oservice.NewExchangeOrder(storages, ticker, userKey, statistics, components)
	bot := bservice.NewBot(storages, ticker, components.Logger, components.Conf)
	webhook := wservice.NewWebhookProcess(storages, order, userKey, bot, components)
	indicator := iservice.NewIndicator(components.Logger, components.Conf, storages)
	client := cservice.NewPlatformBinance(components.RateLimiter)
	clientMeterProxy := emetrics.NewExchangeAPIProxy(client, components.Metrics)
	post := post_service.NewPostService(storages.Post, components.Logger)
	comment := comments_service.NewCommentService(storages.Comment)
	userSub := usubservice.NewUserSubService(storages.UserSub, components.Logger)
	strategy := stservice.NewStrategyService(storages.Strategy, components.Logger, bot)

	services := &Services{
		User:         user,
		Auth:         aservice.NewAuth(user, storages.Verifier, components),
		Bot:          bot,
		ExchangeList: lservice.NewExchangeList(storages.ExchangeList, components),
		Order:        order,
		Ticker:       ticker,
		UserKey:      userKey,
		Webhook:      webhook,
		Indicator:    indicator,
		Client:       clientMeterProxy,
		Post:         post,
		Comments:     comment,
		UserSub:      userSub,
		Strategy:     strategy,
		Statistics:   statistics,
	}
	platform := pservice.NewPlatform(webhook, order, bot, components)
	services.Platform = platform

	return services
}

func NewGRPCServices(storages *storages.Storages, grpcClients *grpc_clients.GRPCClients, components *component.Components) *Services {
	user := uservice.NewUserGRPC(grpcClients.User)
	auth := aservice.NewAuthGRPC(user, grpcClients.Auth, components)
	userKey := kservice.NewUserKey(storages.ExchangeUserKey, components)
	ticker := tiservice.NewTicker(storages.Ticker, components)
	statistics := statservice.NewStatistic(storages.Statistics)
	order := oservice.NewExchangeOrder(storages, ticker, userKey, statistics, components)
	bot := bservice.NewBot(storages, ticker, components.Logger, components.Conf)
	webhook := wservice.NewWebhookProcess(storages, order, userKey, bot, components)
	indicator := iservice.NewIndicator(components.Logger, components.Conf, storages)
	client := cservice.NewPlatformBinance(components.RateLimiter)
	clientMeterProxy := emetrics.NewExchangeAPIProxy(client, components.Metrics)
	post := post_service.NewPostService(storages.Post, components.Logger)
	comment := comments_service.NewCommentService(storages.Comment)
	userSub := usubservice.NewUserSubService(storages.UserSub, components.Logger)
	strategy := stservice.NewStrategyService(storages.Strategy, components.Logger, bot)

	services := &Services{
		User:         user,
		Auth:         auth,
		Bot:          bot,
		ExchangeList: lservice.NewExchangeList(storages.ExchangeList, components),
		Order:        order,
		Ticker:       ticker,
		UserKey:      userKey,
		Webhook:      webhook,
		Indicator:    indicator,
		Client:       clientMeterProxy,
		Post:         post,
		Comments:     comment,
		UserSub:      userSub,
		Strategy:     strategy,
		Statistics:   statistics,
	}
	platform := pservice.NewPlatform(webhook, order, bot, components)
	services.Platform = platform

	return services
}
