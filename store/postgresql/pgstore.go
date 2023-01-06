package postgresql

import (
	"github.com/rs/zerolog/log"

	"github.com/dattranman/todo/model"
	"github.com/dattranman/todo/store"
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	store.Store
	settings model.SQLSettings
	db       *gorm.DB

	task store.TaskStore
}

func NewPostgres(settings model.SQLSettings) *PostgresStore {
	p := &PostgresStore{
		settings: settings,
	}

	p.initConnection()
	p.task = NewTaskStore(p)
	return p
}

func (p *PostgresStore) initConnection() {
	db, err := gorm.Open(p.settings.DriverName, p.settings.URI)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open SQL connection")
	}
	p.db = db
}

func (p *PostgresStore) Task() store.TaskStore {
	return p.task
}
