package postgresql

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dattranman/todo/model/schema"
	"github.com/dattranman/todo/store"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	p    *PostgresStore
	task store.TaskStore
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	p := &PostgresStore{db: s.DB}
	s.task = NewTaskStore(p)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestPostgresTaskStoreCreate() {
	var (
		now = time.Now()
	)
	type fields struct {
		postgres *PostgresStore
	}
	type args struct {
		task *schema.Task
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectQuery string
		wantErr     bool
	}{
		// Test cases.
		{
			name:   "Create success",
			fields: fields{postgres: s.p},
			args: args{
				task: &schema.Task{
					ID:        uuid.New().String(),
					Name:      "test",
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expectQuery: `INSERT INTO "tasks" ("id","name","time","created_at","updated_at","note","deleted_at","point","priority") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			s.mock.ExpectBegin()
			s.mock.ExpectQuery(regexp.QuoteMeta(tt.expectQuery)).
				WithArgs(tt.args.task.ID, tt.args.task.Name, tt.args.task.Time, tt.args.task.CreatedAt, tt.args.task.UpdatedAt, "", nil, 0, 0).
				WillReturnRows(
					sqlmock.NewRows([]string{"id"}).AddRow(tt.args.task.ID),
				)
			s.mock.ExpectCommit()
			err := s.task.Create(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresTaskStore.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
