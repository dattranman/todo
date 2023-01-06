package schema

import (
	"time"

	"github.com/dattranman/todo/model/response"
	"github.com/dattranman/todo/util/json"
)

const (
	taskTable = "tasks"
)

type Task struct {
	ID        string     `json:"id,omitempty" gorm:"column:id;default:uuid_generate_v4()" `
	Name      string     `json:"name" gorm:"column:name" binding:"required,lte=21"`
	Time      time.Time  `json:"time" gorm:"column:time"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	Note      string     `json:"note" gorm:"column:note"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Point     int16      `json:"point" gorm:"column:point"`
	Priority  int8       `json:"priority" gorm:"column:priority"`
}

func (Task) TableName() string {
	return taskTable
}
func (t *Task) ParseToResponse() (task response.Task, err error) {
	byteReq, err := json.Marshal(&t)
	if err != nil {
		return
	}

	err = json.Unmashal(byteReq, &task)
	return
}
