package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GitHubA496/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handleCreateFeedFollowers(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Feed_id uuid.UUID `json:"feed_id"`
	}

	paramas := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&paramas)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}
	// fmt.Print(paramas)

	var feedFollow database.FeedFollower
	feedFollow, err = apiCfg.DB.CreateFeedFollower(r.Context(), database.CreateFeedFollowerParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    paramas.Feed_id,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed, %v", err))
	}

	respondWithJSON(w, 200, databaseFeedFollowtoFeedFollow(feedFollow))
}

func (apiCfg *apiconfig) handleGetFeedFollowers(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollow, err := apiCfg.DB.GetFeedFollowers(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Fetch feed, %v", err))
	}

	respondWithJSON(w, 200, databaseFeedsFollowstoFeedsFollow(feedFollow))
}

func (apiCfg *apiconfig) handleDeleteFeedFollowers(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedFollowIDstr := chi.URLParam(r, "feedFollowId")
	FeedFollowId, err := uuid.Parse(FeedFollowIDstr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Parse the FeedId, %v", err))
	}

	err = apiCfg.DB.DeleteFeedFollower(r.Context(), database.DeleteFeedFollowerParams{
		ID:     FeedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't Delete the Feed, %v", err))
	}
	respondWithJSON(w, 200, struct{}{})
}
