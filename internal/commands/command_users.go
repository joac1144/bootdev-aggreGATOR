package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerUsers(s *state.State, cmd Command) error {

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		log.Fatalf("Error getting all users from database. Error: %s", err)
	}

	for _, user := range users {
		if user.Name == s.Config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
