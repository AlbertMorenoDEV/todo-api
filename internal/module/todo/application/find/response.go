package find

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
	"time"
)

type Response struct {
	id        string    `json:"id"`
	title     string    `json:"title"`
	completed bool      `json:"completed"`
	due       time.Time `json:"due"`
}

func NewResponse(todo todo.Todo) *Response {
	return &Response{
		id:        todo.ID().String(),
		title:     todo.Title().String(),
		completed: todo.Completed().Bool(),
		due:       todo.Due().Time(),
	}
}

func (r Response) ResponseID() string {
	return "find_todo"
}
