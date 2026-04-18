package handler

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (h * Handler) Router() chi.Router {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: h.allowedOrigins,
		AllowedMethods: []string{"GET" , "POST" , "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Post("/bundles" , h.CreateBundle)
	r.Get("/bundles/{slug}" , h.GetBundle)

	return r

}