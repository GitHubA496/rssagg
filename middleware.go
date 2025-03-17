package main

import (
	"fmt"
	"net/http"

	"github.com/GitHubA496/rssagg/internal/auth"
	"github.com/GitHubA496/rssagg/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiconfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Countn't get a user %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("Countn't get a user %v", err))
			return
		}
		handler(w, r, user)
	}
}
