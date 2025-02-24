package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	return router
}
