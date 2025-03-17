package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GitHubA496/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name    string `json:"name"`
		Url     string `json:"url"`
		User_id string `json:"user_id"`
	}

	paramas := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&paramas)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}
	// fmt.Print(paramas)

	var feed database.Feed
	feed, err = apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      paramas.Name,
		Url:       paramas.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed, %v", err))
	}

	respondWithJSON(w, 200, databaseFeedtoFeed(feed))
}

func (apiCfg *apiconfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {
	var feeds []database.Feed
	feeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds, %v", err))
	}
	respondWithJSON(w, 200, databaseFeedstoFeeds(feeds))
}
