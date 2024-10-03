package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/metrics"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/middleware"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules"
)

func NewApiRouter(controllers *modules.Controllers, components *component.Components) http.Handler {
	r := chi.NewRouter()
	//for prometheus
	r.Get("/metrics", promhttp.Handler().ServeHTTP)
	r.Route("/echarts", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			metrics.InitHTTPHandler().ServeHTTP(w, r)
		})
	})
	//platformController := controllers.Platformer
	r.Route("/api", func(r chi.Router) {
		r.Route("/1", func(r chi.Router) {
			//for testing grpc microservice
			authCheck := middleware.NewTokenManager(components.Responder, components.TokenManager)
			r.Route("/auth", func(r chi.Router) {
				r.Post("/register", controllers.Auth.Register)
				r.Post("/login", controllers.Auth.Login)
				r.Route("/refresh", func(r chi.Router) {
					r.Use(authCheck.CheckRefresh)
					r.Post("/", controllers.Auth.Refresh)
				})
				r.Route("/{provider}", func(r chi.Router) {
					r.Get("/login", controllers.Auth.SocialRedirect)
					r.Get("/callback", controllers.Auth.SocialCallback)
				})
				r.Post("/verify", controllers.Auth.Verify)
			})
			r.Route("/user", func(r chi.Router) {
				{
					//r.Use(authCheck.CheckStrict)
					r.Post("/changePassword", controllers.User.ChangePassword)
					r.Post("/profile", controllers.User.Profile)
					r.Post("/sendCode", controllers.User.SendResetCodeEmail)
					r.Post("/resetPassword", controllers.User.ResetPassword)

				}

				r.Route("/subscription", func(r chi.Router) {
					r.Post("/add", controllers.UserSub.Add)
					r.Get("/list", controllers.UserSub.List)
					r.Delete("/delete", controllers.UserSub.Delete)

				})
			})

			r.Route("/posts", func(r chi.Router) {
				r.Use(authCheck.CheckStrict)
				r.Post("/create", controllers.Post.CreatePost)
				r.Post("/delete", controllers.Post.DeletePost)
				r.Post("/update", controllers.Post.UpdatePost)
				r.Get("/id", controllers.Post.GetByIdUser)
				r.Get("/tape", controllers.Post.GetListTape)
			})

			r.Route("/comments", func(r chi.Router) {
				r.Use(authCheck.CheckStrict)
				commentsController := controllers.Comments
				r.Post("/create", commentsController.CreateComment)
				r.Post("/delete", commentsController.DeleteComment)
				r.Post("/update", commentsController.UpdateComment)
				r.Get("/id", commentsController.GetCommentByID)
				r.Get("/list", commentsController.GetCommentsList)
			})

			r.Route("/strategy", func(r chi.Router) {
				r.Use(authCheck.CheckStrict)
				r.Get("/create", controllers.Bot.Create)
				r.Post("/update", controllers.Bot.Update)
				r.Post("/get", controllers.Bot.Get)
				r.Get("/list", controllers.Bot.List)
				r.Post("/toggle", controllers.Bot.Toggle)
				r.Post("/delete", controllers.Bot.Delete)
				r.Post("/signals", controllers.Bot.WebhookSignal)
			})

			r.Route("/bot", func(r chi.Router) {
				r.Use(authCheck.CheckStrict)
				r.Get("/create", controllers.Bot.Create)
				r.Post("/update", controllers.Bot.Update)
				r.Post("/get", controllers.Bot.Get)
				r.Get("/list", controllers.Bot.List)
				r.Post("/toggle", controllers.Bot.Toggle)
				r.Post("/delete", controllers.Bot.Delete)
				r.Post("/signals", controllers.Bot.WebhookSignal)
			})

			r.Route("/exchange", func(r chi.Router) {
				r.Use(authCheck.CheckStrict)
				r.Get("/ticker", controllers.Ticker.Ticker)

				r.Post("/add", controllers.List.ExchangeAdd)
				r.Post("/delete", controllers.List.ExchangeDelete)
				r.Get("/list", controllers.List.ExchangeList)

				r.Get("/pairs", controllers.Indicator.GetDynamicPairs)
				r.Post("/ema", controllers.Indicator.EMA)
				r.Route("/balance", func(r chi.Router) {
					r.Use(authCheck.CheckStrict)
					r.Get("/", controllers.Client.GetBalanceAccount)
				})
				r.Get("/candles", controllers.Client.GetCandles)

				r.Route("/user", func(r chi.Router) {
					r.Route("/key", func(r chi.Router) {
						r.Use(authCheck.CheckStrict)
						r.Post("/add", controllers.UserKey.ExchangeUserKeyAdd)
						r.Post("/delete", controllers.UserKey.ExchangeUserKeyDelete)
						r.Get("/list", controllers.UserKey.ExchangeUserKeyList)
					})
					r.Route("/orders", func(r chi.Router) {
						r.Use(authCheck.CheckStrict)
						r.Get("/history", controllers.Order.UserOrdersHistory)
					})
					r.Route("/webhooks", func(r chi.Router) {
						r.Use(authCheck.CheckStrict)
						r.Get("/history", controllers.Webhook.UserWebhooksHistory)
						r.Post("/toggle", controllers.Webhook.UserWebhooksHistory)
						r.Post("/info", controllers.Webhook.WebhookInfo)
					})
					r.Route("/bot", func(r chi.Router) {
						r.Use(authCheck.CheckStrict)
						r.Post("/list", controllers.Order.GetBotOrders)
						r.Post("/info", controllers.Webhook.BotInfo)
					})
				})
			})
			r.Route("/statistics", func(r chi.Router) {
				r.Use(authCheck.CheckStrict)
				r.Get("/user", controllers.Statistic.GetUserStatistic)
				r.Post("/bot", controllers.Statistic.GetBotStatistic)
				r.Post("/delete", controllers.Statistic.DeleteStatistic)

			})
		})
	})

	r.Route("/platform", func(r chi.Router) {
		r.Post("/hook", controllers.Webhook.WebhookProcess)
	})

	return r
}
