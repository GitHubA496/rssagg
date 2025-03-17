package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/GitHubA496/rssagg/internal/database"
	"github.com/google/uuid"
)

func startScrapping(db *database.Queries, concurreny int, timeBetweenRequest time.Duration) {
	log.Printf("Scrapping started with %d concurreny and %s between requests", concurreny, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurreny))
		if err != nil {
			log.Fatal("unble to fetch Next feed", err)
			continue
		}
		wg := sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, &wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println(err)
		return
	}
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error while fetching feed", err)
		return
	}
	for _, item := range rssFeed.Channel.Items {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		t, err := time.Parse(time.RFC1123, item.Pubdate)
		if err != nil {
			log.Println("Error while parsing time", err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: description,
			PublishedAt: t,
			Url:         item.Link,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Error while creating post", err)
			continue
		}
	}
	log.Println("Feed", feed.Name, "has", len(rssFeed.Channel.Items), "items")

}
