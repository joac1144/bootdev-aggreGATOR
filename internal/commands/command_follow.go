package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerFollow(s *state.State, cmd Command, currentUser database.User) error {
	url := cmd.Args[0]

	ctx := context.Background()

	feed, err := s.Db.GetFeedByUrl(ctx, url)
	if err != nil {
		return err
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}
	feed_follow, err := s.Db.CreateFeedFollow(ctx, createFeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed name: %s, current user: %s\n", feed_follow.FeedName, feed_follow.UserName)

	return nil
}
