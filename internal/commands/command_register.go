package commands

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

func HandlerRegister(s *state.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("command 'register' requires 1 argument")
	}

	username := cmd.Args[0]

	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	user, err := s.Db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		log.Fatalf("Failed creating user \"%s\". User may already exist.", username)
	}

	s.Config.SetUser(username)

	fmt.Printf("User \"%s\" has been created\n", username)
	log.Println("Created User:")
	fmt.Printf("%+v\n", user)

	return nil
}
