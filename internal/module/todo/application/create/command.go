package create

type Command struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Due   int64  `json:"due"`
}

func (c Command) CommandID() string {
	return "create_todo"
}
