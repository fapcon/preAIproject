package modules

import (
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	acontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/auth/controller"
	scontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/bot/controller"
	comments_controller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/comments/controller"
	ccontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/controller"
	icontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_indicator/controller"
	lcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_list/controller"
	ocontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_order/controller"
	ticontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_ticker/controller"
	kcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_user_key/controller"
	pcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/platform/controller"
	post_controller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/posts/controller"
	statcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/statistics/controller"
	stcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/strategy/controller"
	tcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/testGRPC/controller"
	ucontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user/controller"
	usubcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/user_sub/controller"
	wcontroller "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/webhook_process/controller"
)

type Controllers struct {
	Auth       acontroller.Auther
	User       ucontroller.Userer
	Bot        scontroller.Boter
	List       lcontroller.ExchangeLister
	Order      ocontroller.ExchangeOrderer
	Ticker     ticontroller.ExchangeTicker
	UserKey    kcontroller.ExchangeUserKeyer
	Platformer pcontroller.Platformer
	TestGrpc   tcontroller.GrpcTester
	Webhook    wcontroller.WebhookProcesser
	Indicator  icontroller.Indicatorer
	Client     ccontroller.Exchanger
	Post       post_controller.Poster
	Comments   comments_controller.Comenter
	UserSub    usubcontroller.UserSuber
	Strategy   stcontroller.Strateger
	Statistic  statcontroller.Statisticer
}

func NewControllers(services *Services, components *component.Components) *Controllers {
	authController := acontroller.NewAuth(services.Auth, components)
	userController := ucontroller.NewUser(services.User, components)
	botController := scontroller.NewBot(services.Bot, components)

	return &Controllers{
		Auth:       authController,
		User:       userController,
		Bot:        botController,
		List:       lcontroller.NewExchangeList(services.ExchangeList, components),
		Order:      ocontroller.NewExchangeOrder(services.Order, components),
		Ticker:     ticontroller.NewTicker(services.Ticker, components),
		UserKey:    kcontroller.NewUserKey(services.UserKey, components),
		Platformer: pcontroller.NewPlatform(services.Platform, components),
		Webhook:    wcontroller.NewWebhookProcess(services.Webhook, components),
		Indicator:  icontroller.NewIndicator(services.Indicator, components),
		Client:     ccontroller.NewClient(services.Client, components),
		Post:       post_controller.NewPost(services.Post, components),
		Comments:   comments_controller.NewCommentController(services.Comments, components),
		UserSub:    usubcontroller.NewUserSubscription(services.UserSub, components),
		Strategy:   stcontroller.NewStrategyController(services.Strategy, components),
		Statistic:  statcontroller.NewStatisticController(services.Statistics, components),
	}
}
