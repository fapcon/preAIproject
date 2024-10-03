package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/module/auth"
	"studentgit.kata.academy/eazzyearn/students/mono/auth-sso/server/handlers"
	"time"
)

func CreateHTTPServer(httpAddr string, authHandler handlers.Auther) *http.Server {
	r := chi.NewRouter()

	r.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: true,
				MaxAge:           300,
			},
		),
	)

	r.Route("/api/v1/user", func(r chi.Router) {
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)
		r.Get("/profile", authHandler.Profile)
		r.Get("/logout", authHandler.Logout)

		r.Route("/{provider}", func(r chi.Router) {
			r.Get("/login", authHandler.SocialGetRedirectURL)
			r.Get("/callback", authHandler.SocialCallback)
		})
	})

	r.Get("/docs", auth.SwaggerUI)
	r.Get("/swagger/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger"))).ServeHTTP(w, r)
	})

	return &http.Server{
		Addr:              httpAddr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
}
