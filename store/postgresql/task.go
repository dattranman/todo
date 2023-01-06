package postgresql

import (
	"github.com/dattranman/todo/model/schema"
	"github.com/dattranman/todo/store"
	"github.com/dattranman/todo/util/json"
)

type PostgresTaskStore struct {
	postgres *PostgresStore
}

func NewTaskStore(pg *PostgresStore) store.TaskStore {
	return &PostgresTaskStore{
		postgres: pg,
	}
}

func (p *PostgresTaskStore) Create(task *schema.Task) (err error) {
	err = p.postgres.db.Create(task).Error
	return
}

func (p *PostgresTaskStore) GetByID(id string) (task schema.Task, err error) {
	err = p.postgres.db.Where(&schema.Task{ID: id}).
		First(&task).Error
	return
}

func (p *PostgresTaskStore) GetByIDs(ids []string) (list []schema.Task, err error) {
	return
}

func (p *PostgresTaskStore) GetList() (list []schema.Task, total int64, err error) {
	err = p.postgres.db.Model(&schema.Task{}).Find(&list).Error
	return
}

func (p *PostgresTaskStore) Update(task *schema.Task) (err error) {
	byteData, err := json.Marshal(task)
	if err != nil {
		return
	}
	mapUpdate := make(map[string]interface{})
	err = json.Unmashal(byteData, &mapUpdate)
	if err != nil {
		return
	}
	err = p.postgres.db.Model(&schema.Task{}).Where("id = ?", task.ID).Updates(mapUpdate).Error
	return
}

func (p *PostgresTaskStore) Delete(id string) (err error) {
	err = p.postgres.db.Where(&schema.Task{ID: id}).
		Delete(&schema.Task{}).Error
	return
}
