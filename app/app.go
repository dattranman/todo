package app

import (
	"github.com/dattranman/todo/config"
	"github.com/dattranman/todo/model"
	"github.com/dattranman/todo/store"
	"github.com/dattranman/todo/store/postgresql"

	"github.com/rs/zerolog/log"
)

type App struct {
	Config model.Configuration

	Store store.Store
}

func New(cfg string) (*App, error) {
	c, err := config.Load(cfg)
	if err != nil {
		return nil, err
	}
	app := &App{
		Config: *c,
	}

	log.Info().Msg("Server is initializing...")
	app.Store = postgresql.NewPostgres(app.Config.SQLSettings)
	if err != nil {
		return nil, err

	}

	return app, nil
}
