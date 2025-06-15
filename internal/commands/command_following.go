package commands

import (
	"context"
	"fmt"

	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerFollowing(s *state.State, cmd Command, currentUser database.User) error {
	feed_follows, err := s.Db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return nil
	}

	fmt.Println("You follow these feeds:")
	for _, feed_follow := range feed_follows {
		fmt.Println(feed_follow.FeedName)
	}

	return nil
}
