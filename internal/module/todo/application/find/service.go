package find

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
)

// Service that allow finding todos
type Service interface {
	FindById(id identifier.Identifier) (*todo.Todo, error)
}

type service struct {
	repository todo.Repository
}

// NewService creates a finder service with the necessary dependencies
func NewService(repository todo.Repository) Service {
	return &service{repository}
}

// CreateTodo creates a new item on the repository
func (s *service) FindById(id identifier.Identifier) (*todo.Todo, error) {
	return s.repository.Find(id)
}
