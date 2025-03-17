package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/GitHubA496/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	paramas := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&paramas)
	if err != nil {
		respondWithError(w, 400, "Invalid request payload")
		return
	}

	var user database.User
	user, err = apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      paramas.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Countm't create a user %v", err))
	}

	respondWithJSON(w, 200, databaseUsertoUser(user))
}

func (apiCfg *apiconfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, databaseUsertoUser(user))
}

func (apiCfg *apiconfig) handleGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts, %v", err))
	}
	respondWithJSON(w, 200, databasePoststoPosts(posts))
}
