package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	rss "github.com/joac1144/bootdev-aggreGATOR/RSS"
	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerAgg(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("usage: agg <time_between_reqs>")
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return errors.New("failed to parse input as time")
	}

	fmt.Printf("Collecting feeds every %s...\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state.State) {
	nextFeed, err := s.Db.GetNextFeedFetch(context.Background())
	if err != nil {
		log.Printf("Error fetching next feed: %s", err)
		return
	}

	err = s.Db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %s", err)
		return
	}

	rssFeed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s: %s", nextFeed.Name, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		parsedPublishedTime, err := parsePublishedTime(item.PubDate)
		pubTimeIsValid := true
		if err != nil {
			log.Printf("Failed to parse published time for post '%s': %v", item.Title, err)
			pubTimeIsValid = false
		}

		timeNow := time.Now()
		createPostParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   timeNow,
			UpdatedAt:   timeNow,
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: parsedPublishedTime, Valid: pubTimeIsValid},
			FeedID:      nextFeed.ID,
		}
		_, err = s.Db.CreatePost(context.Background(), createPostParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Error inserting post '%s': %v", item.Title, err)
		}
		fmt.Printf("Inserted post: %s\n", item.Title)
	}
}

func parsePublishedTime(pubDate string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC3339,
		"Mon, 02 Jan 2006 15:04:05 MST",
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, pubDate); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unable to parse time: %s", pubDate)
}
