package commands

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("command 'login' requires 1 argument")
	}

	username := cmd.Args[0]

	user, err := s.Db.GetUserByName(context.Background(), username)
	if err != nil {
		log.Fatalf("Error getting the user \"%s\". User may not exist in the database.", username)
	}

	if err := s.Config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User has been set to \"%s\"\n", user.Name)

	return nil
}
