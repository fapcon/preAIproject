package middleware

import (
	"github.com/go-chi/cors"
	"net/http"
	"studentgit.kata.academy/eazzyearn/students/mono/monolith/config"
)

type Cors struct {
	conf config.Cors
}

func NewCors() *Cors {
	return &Cors{}
}

func (t *Cors) OpenAllCors(next http.Handler) http.Handler {
	return cors.Handler(
		cors.Options{
			AllowedOrigins:   t.conf.AllowedOrigins,
			AllowedMethods:   t.conf.AllowedMethods,
			AllowedHeaders:   t.conf.AllowedHeaders,
			ExposedHeaders:   t.conf.ExposedHeaders,
			AllowCredentials: t.conf.AllowCredentials,
			MaxAge:           t.conf.MaxAge, // Maximum value not ignored by any of major browsers
		},
	)(next)
}
