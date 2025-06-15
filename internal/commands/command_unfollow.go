package commands

import (
	"context"
	"fmt"

	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerUnfollow(s *state.State, cmd Command, user database.User) error {
	url := cmd.Args[0]

	deleteFeedFollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	}

	err := s.Db.DeleteFeedFollow(context.Background(), deleteFeedFollowParams)
	if err != nil {
		return err
	}

	fmt.Printf("Unfollowed %s as user %s", url, user.Name)

	return nil
}
