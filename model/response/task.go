package response

import "time"

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Note      string    `json:"note"`
	Point     int16     `json:"point"`
	Priority  int8      `json:"priority"`
	Time      time.Time `json:"time"`
	CreatedAt time.Time `json:"created_at"`
}

type ListTasks struct {
	List []Task `json:"list"`
}
