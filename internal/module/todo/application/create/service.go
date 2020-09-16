package create

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/due"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/title"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
)

// Service that allow creating todos
type Service interface {
	CreateTodo(id identifier.Identifier, title title.Title, due due.Due) error
}

type service struct {
	repository todo.Repository
}

// NewService creates a creator service with the necessary dependencies
func NewService(repository todo.Repository) Service {
	return &service{repository}
}

// CreateTodo creates a new item on the repository
func (s *service) CreateTodo(id identifier.Identifier, title title.Title, due due.Due) error {
	newTodo, err := todo.NewTodo(id, title, due)
	if err != nil {
		return err
	}

	return s.repository.Save(newTodo)
}
