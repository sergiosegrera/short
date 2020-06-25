package handlers

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sergiosegrera/short/service"
)

func MakeGetLinkHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		response, err := svc.GetLink(r.Context(), id)
		if err != nil {
			switch err {
			case service.ErrIdNotFound:
				JSON(w, 500, message{"error": "Id not found"})
				return
			default:
				JSON(w, 500, message{"error": "An unknown error occured"})
				return
			}
		}

		JSON(w, 200, response)
	}
}
