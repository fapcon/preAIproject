package run

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/grpc_clients"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/metrics"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/ptflp/godecoder"
	"gitlab.com/golight/orm"
	"gitlab.com/golight/orm/utils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tg_bot"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules/exchange_client/service"

	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/cache"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/errors"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/responder"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/router"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/server"
	internal "studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/service"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/tools/cryptography"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/models"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/provider"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/storages"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/worker"
)

type Application interface {
	Runner
	Bootstraper
}

type Runner interface {
	Run() int
}

type Bootstraper interface {
	Bootstrap(options ...interface{}) Runner
}

type App struct {
	conf     config.AppConf
	logger   *zap.Logger
	srv      server.Server
	Sig      chan os.Signal
	Storages *storages.Storages
	Servises *modules.Services
}

func NewApp(conf config.AppConf, logger *zap.Logger) *App {
	return &App{conf: conf, logger: logger, Sig: make(chan os.Signal, 1)}
}

// на русском
// Bootstrap - инициализация приложения
// options - дополнительные параметры
// возвращает интерфейс Runner
func (a *App) Run() int {
	// shutdown server on signal interrupt
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigInt := <-a.Sig
		a.logger.Info("signal interrupt received", zap.Stringer("os_signal", sigInt))
		cancel()
	}()

	errGroup, ctx := errgroup.WithContext(context.Background())
	errGroup.Go(func() error {
		err := a.srv.Serve(ctx)
		if err != nil && err != http.ErrServerClosed {
			a.logger.Error("app: server error", zap.Error(err))
			return err
		}
		return nil
	})

	if err := errGroup.Wait(); err != nil {
		return errors.GeneralError
	}

	return errors.NoError
}

// на русском
// Bootstrap - инициализация приложения
func (a *App) Bootstrap(options ...interface{}) Runner {
	email := provider.NewEmail(a.conf.Provider.Email, a.logger)
	notifyEmail := internal.NewNotify(a.conf.Provider.Email, email, a.logger)
	tokenManager := cryptography.NewTokenJWT(a.conf.Token)
	decoder := godecoder.NewDecoder(jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		DisallowUnknownFields:  true,
	})
	registry := prometheus.NewRegistry()
	prometheusMetrics := metrics.NewPrometheusMetrics(registry)
	responseManager := responder.NewResponder(decoder, a.logger, prometheusMetrics)
	uuID := cryptography.NewUUIDGenerator()
	hash := cryptography.NewHash(uuID)
	errChan := make(chan models.ErrMessage)
	notificationChan := make(chan models.NotificationMessage)
	components := component.NewComponents(a.conf, notifyEmail, tokenManager, responseManager, decoder, hash, a.logger, errChan, prometheusMetrics)

	tableScanner := utils.NewTableScanner()
	tableScanner.RegisterTable(
		&models.UserDTO{},
		&models.BotDTO{},
		&models.ExchangeListDTO{},
		&models.ExchangeTickerDTO{},
		&models.ExchangeUserKeyDTO{},
		&models.ExchangeOrderDTO{},
		&models.ExchangeOrderLogDTO{},
		&models.ExchangeOrderDTO{},
		&models.StrategyPairDTO{},
		&models.WebhookProcessDTO{},
		&models.WebhookProcessHistoryDTO{},
		&models.EmailVerifyDTO{},
		&models.StrategyDTO{},
		&models.StrategySubscribersDTO{},
		&models.PostDTO{},
		&models.UsersBanedDTO{},
		&models.CommentDTO{},
		&models.UserSubDTO{},
		&models.BotStatisticsDTO{},
	)
	sqlAdapter, err := orm.NewOrm(utils.DB(a.conf.DB), tableScanner, a.logger)
	if err != nil {
		a.logger.Fatal("error init db", zap.Error(err))
	}
	cacheClient, err := cache.NewCache(a.conf.Cache, decoder, a.logger)
	if err != nil {
		a.logger.Fatal("error init cache", zap.Error(err))
	}
	newStorages := storages.NewStorages(sqlAdapter, cacheClient)
	a.Storages = newStorages

	telegramBot := tg_bot.NewTgBotClient(a.conf, a.logger)
	go func() {
		telegramBot.Run(components.ErrChan)
	}()

	var services *modules.Services

	switch a.conf.GRPC.Type {
	case "grpc":
		conn := a.initGrpcClient(a.conf.GRPC.Host, a.conf.GRPC.Port)
		clients := grpc_clients.NewGRPCClients(conn)
		services = modules.NewGRPCServices(newStorages, clients, components)
		a.Servises = services

	case "local":
		services = modules.NewServices(newStorages, components)
		a.Servises = services
	}

	tickerWorker := worker.NewTicker(service.NewPlatformBinance(components.RateLimiter), *services, service.Binance, notificationChan)
	go func() {
		telegramBot.RunNotification(notificationChan)
		err = tickerWorker.Run()
		if err != nil {
			a.logger.Error("ticker worker error", zap.Error(err))
		}
	}()
	orderStatus := worker.NewOrderStatus(newStorages, *services, components)
	go func() {
		err = orderStatus.Run()
		if err != nil {
			a.logger.Error("ticker worker error", zap.Error(err))
		}
	}()

	// инициализация контроллеров
	controllers := modules.NewControllers(services, components)

	// инициализация роутера

	r := router.NewRouter(controllers, components)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.conf.Server.Port),
		Handler: r,
	}

	// инициализация сервера
	a.srv = server.NewHTTPServer(a.conf.Server, srv, a.logger, registry)

	return a
}

func (a *App) initGrpcClient(host, port string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		a.logger.Error(fmt.Sprintf("Error connecting grpc client %s:", host), zap.Error(err))
		return nil
	} else {
		a.logger.Info(fmt.Sprintf("rpc client %s connected", host))
	}
	return client
}
