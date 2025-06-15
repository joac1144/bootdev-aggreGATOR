package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerFeeds(s *state.State, cmd Command) error {
	ctx := context.Background()
	feeds, err := s.Db.GetFeeds(ctx)
	if err != nil {
		log.Fatalf("Error getting all feeds from database. Error: %s", err)
	}

	for _, feed := range feeds {
		user, err := s.Db.GetUserById(ctx, feed.UserID)
		if err != nil {
			log.Fatalf("Error getting user with id %s from database. Error: %s", feed.UserID, err)
		}

		fmt.Printf("Name: %s, URL: %s, Created by: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
