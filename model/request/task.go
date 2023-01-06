package request

import (
	"time"

	"github.com/dattranman/todo/model/schema"
	"github.com/dattranman/todo/util/json"
)

type Task struct {
	Name     string    `json:"name"`
	Time     time.Time `json:"time"`
	Point    int16     `json:"point"`
	Priority int8      `json:"priority"`
}

func (t *Task) ParseToSchema() (task schema.Task, err error) {
	byteReq, err := json.Marshal(t)
	if err != nil {
		return
	}

	err = json.Unmashal(byteReq, &task)
	return
}
