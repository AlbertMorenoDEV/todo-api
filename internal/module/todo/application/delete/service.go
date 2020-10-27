package delete

import (
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/todo"
)

// Service that allow deleting todos
type Service interface {
	DeleteTodo(id identifier.Identifier) error
}

type service struct {
	repository todo.Repository
}

// NewService creates a deleter service with the necessary dependencies
func NewService(repository todo.Repository) Service {
	return &service{repository}
}

// DeleteTodo deletes a item on the repository
func (s *service) DeleteTodo(id identifier.Identifier) error {
	return s.repository.Delete(id)
}
