package delete

type Command struct {
	ID string `json:"id"`
}

func (c Command) CommandID() string {
	return "delete_todo"
}
