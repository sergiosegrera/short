package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sergiosegrera/short/service"
)

type CreateLinkRequest struct {
	Url string `json:"url"`
}

func MakeCreateLinkHandler(svc service.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: No need for id here we shouldnt ask for a "Link" object
		var request CreateLinkRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			JSON(w, 400, message{"error": "Invalid request, failed to decode url"})
			return
		}

		response, err := svc.CreateLink(r.Context(), request.Url)
		if err != nil {
			switch err {
			case service.ErrCreatingLink:
				JSON(w, 500, message{"error": "Error creating link"})
				return
			default:
				JSON(w, 500, message{"error": "An unknown error occured"})
				return
			}
		}

		JSON(w, 200, response)
	}
}
