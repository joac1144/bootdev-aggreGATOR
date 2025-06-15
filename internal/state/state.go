package state

import (
	"github.com/joac1144/bootdev-aggreGATOR/internal/config"
	"github.com/joac1144/bootdev-aggreGATOR/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}
