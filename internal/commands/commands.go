package commands

import (
	"context"
	"errors"

	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Commands map[string]func(*state.State, Command) error
}

func (cmds *Commands) Run(s *state.State, cmd Command) error {
	fun, ok := cmds.Commands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return fun(s, cmd)
}

func (cmds *Commands) Register(name string, f func(*state.State, Command) error) {
	cmds.Commands[name] = f
}

func MiddlewareLoggedIn(handler func(s *state.State, cmd Command, user database.User) error) func(*state.State, Command) error {
	return func(s *state.State, c Command) error {
		user, err := s.Db.GetUserByName(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
