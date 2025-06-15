package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerReset(s *state.State, cmd Command) error {

	err := s.Db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatalf("Error deleting all users from database. Error: %s", err)
	}

	fmt.Println("Reset all users in database")

	return nil
}
