package find

type Query struct {
	ID string `json:"id"`
}

func (q Query) QueryID() string {
	return "found_todo"
}
