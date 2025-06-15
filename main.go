package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joac1144/bootdev-aggreGATOR/internal/commands"
	"github.com/joac1144/bootdev-aggreGATOR/internal/config"
	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
	"github.com/joac1144/bootdev-aggreGATOR/internal/state"
	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", conf.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	s := &state.State{
		Db:     dbQueries,
		Config: &conf,
	}

	cmds := commands.Commands{
		Commands: make(map[string]func(*state.State, commands.Command) error),
	}

	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerUsers)
	cmds.Register("agg", commands.HandlerAgg)
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("feeds", commands.HandlerFeeds)
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerUnfollow))

	args := os.Args
	if len(args) < 2 {
		log.Fatal("You must provide a command")
	}

	input_cmd := args[1]
	cmd := commands.Command{
		Name: input_cmd,
		Args: args[2:],
	}

	err = cmds.Run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
