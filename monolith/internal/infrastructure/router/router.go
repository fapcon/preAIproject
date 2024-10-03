package router

import (
	chi "github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/component"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/infrastructure/middleware"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/modules"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/internal/router"
)

func NewRouter(controllers *modules.Controllers, components *component.Components) *chi.Mux {
	r := chi.NewRouter()
	setBasicMiddlewares(r)
	setDefaultRoutes(r)

	r.Mount("/", router.NewApiRouter(controllers, components))
	return r
}

func setBasicMiddlewares(r *chi.Mux) {
	proxy := middleware.NewReverseProxy()
	//corsHeader := middleware.NewCors()
	r.Use(chimw.Recoverer)
	r.Use(proxy.ReverseProxy)
	r.Use(
		cors.Handler(
			cors.Options{
				// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
				AllowedOrigins: []string{"https://*", "http://*"},
				// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300, // Maximum value not ignored by any of major browsers
			},
		),
	)
	//r.Use(corsHeader.OpenAllCors)
	r.Use(chimw.RealIP)
	r.Use(chimw.RequestID)
}

func setDefaultRoutes(r *chi.Mux) {
	r.Get("/swagger", swaggerUI)
	r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP(w, r)
	})
}
