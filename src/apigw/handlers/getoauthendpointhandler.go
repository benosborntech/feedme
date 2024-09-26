package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benosborntech/feedme/apigw/oauth"
)

type handlerResponse struct {
	RedirectEndpoint string `json:"redirectEndpoint"`
	Session          string `json:"session"`
}

func GetOAuthEndpointHandler(oauth oauth.OAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, endpoint, err := oauth.GetEndpoint(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get redirect endpoint, err=%v", err), http.StatusInternalServerError)
			return
		}

		// Prepare the response
		response := &handlerResponse{
			RedirectEndpoint: endpoint,
			Session:          session,
		}

		// Marshal the response to JSON
		payload, err := json.Marshal(response)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to marshal response, err=%v", err), http.StatusInternalServerError)
			return
		}

		// Send the JSON response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}
