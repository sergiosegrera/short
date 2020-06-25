package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/short/service"
	"github.com/sergiosegrera/short/transports/http/handlers"
)

func Serve(svc service.Service) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))

	router.Post("/create", handlers.MakeCreateLinkHandler(svc))
	router.Get("/{id}", handlers.MakeGetLinkHandler(svc))

	return http.ListenAndServe(fmt.Sprintf(":%v", "8080"), router)
}
