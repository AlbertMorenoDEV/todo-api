package create

import "time"

type Command struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Due   time.Time `json:"due"`
}

func (c Command) CommandID() string {
	return "create_todo"
}
