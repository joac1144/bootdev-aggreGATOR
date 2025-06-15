package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerAddFeed(s *state.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("usage: addfeed <feed name> <feed url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	timeNow := time.Now()

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.Db.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return err
	}

	createFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.Db.CreateFeedFollow(context.Background(), createFeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("Created feed %s and started following it\n", feed.Name)
	fmt.Printf("%+v\n", feed)

	return nil
}
