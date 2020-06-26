package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sergiosegrera/short/config"
	"github.com/sergiosegrera/short/service"
	"github.com/sergiosegrera/short/transports/http/handlers"
)

func Serve(svc service.Service, conf *config.Config) error {
	router := chi.NewRouter()
	router.Use(middleware.Compress(5, "gzip"))

	// TODO: Better way to create handlers?
	router.Post("/create", handlers.MakeCreateLinkHandler(svc))
	router.Get("/{id}", handlers.MakeGetLinkHandler(svc))

	return http.ListenAndServe(fmt.Sprintf(":%v", conf.HttpPort), router)
}
