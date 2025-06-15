package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerBrowse(s *state.State, cmd Command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.Args) == 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
	}

	getPostsForUserPrams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.Db.GetPostsForUser(context.Background(), getPostsForUserPrams)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.Title)
	}

	return nil
}
