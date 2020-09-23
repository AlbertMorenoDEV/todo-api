package find

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/query"
)

type Response struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Due       int64  `json:"due"`
}

func NewResponse(todo todo.Todo) query.Response {
	return Response{
		ID:        todo.ID().String(),
		Title:     todo.Title().String(),
		Completed: todo.Completed().Bool(),
		Due:       todo.Due().Time().Unix(),
	}
}

func (r Response) ResponseID() string {
	return "find_todo"
}
