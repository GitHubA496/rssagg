package main

import (
	"database/sql"
	"time"

	"github.com/GitHubA496/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upadated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upadated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.CreatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseFeedtoFeed(dbUser database.Feed) Feed {
	return Feed{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.CreatedAt,
		Name:      dbUser.Name,
		Url:       dbUser.Url,
		UserID:    dbUser.UserID,
	}
}
func databaseFeedstoFeeds(dbUser []database.Feed) []Feed {
	var feeds []Feed
	for _, feed := range dbUser {
		feeds = append(feeds, databaseFeedtoFeed(feed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"upadated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowtoFeedFollow(dbUser database.FeedFollower) FeedFollow {
	return FeedFollow{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.CreatedAt,
		UserID:    dbUser.UserID,
		FeedID:    dbUser.FeedID,
	}
}
func databaseFeedsFollowstoFeedsFollow(dbUser []database.FeedFollower) []FeedFollow {
	var feeds []FeedFollow
	for _, feedFollow := range dbUser {
		feeds = append(feeds, databaseFeedFollowtoFeedFollow(feedFollow))
	}
	return feeds
}

type Post struct {
	ID          uuid.UUID      `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"upadated_at"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	PublishedAt time.Time      `json:"published_at"`
	Url         string         `json:"url"`
	FeedID      uuid.UUID      `json:"feed_id"`
}

func databasePosttoPost(dbUser database.Post) Post {
	return Post{
		ID:          dbUser.ID,
		CreatedAt:   dbUser.CreatedAt,
		UpdatedAt:   dbUser.CreatedAt,
		Title:       dbUser.Title,
		Description: dbUser.Description,
		PublishedAt: dbUser.PublishedAt,
		Url:         dbUser.Url,
		FeedID:      dbUser.FeedID,
	}
}

func databasePoststoPosts(dbUser []database.Post) []Post {
	var posts []Post
	for _, post := range dbUser {
		posts = append(posts, databasePosttoPost(post))
	}
	return posts
}
