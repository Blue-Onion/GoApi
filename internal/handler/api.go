package handler

import (
	"github.com/blue-onion/goApi/internal/middeware"
	"github.com/go-chi/chi"
	chiMiddle "github.com/go-chi/chi/middleware"
)
func Handler(r *chi.Mux){
	// Global Middleware
	r.Use(chiMiddle.StripSlashes)
	r.Route("/account",func(r chi.Router) {
		r.Use(middeware.Authorization)
		r.Get("/coins",GetCoinBalance)
	})
}