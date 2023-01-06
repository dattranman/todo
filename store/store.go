package store

import (
	"github.com/dattranman/todo/model/schema"
)

type Store interface {
	Task() TaskStore
}

type TaskStore interface {
	Create(task *schema.Task) (err error)
	GetByID(id string) (task schema.Task, err error)
	GetByIDs(ids []string) (list []schema.Task, err error)
	GetList() (list []schema.Task, total int64, err error)
	Update(task *schema.Task) (err error)
	Delete(id string) (err error)
}
