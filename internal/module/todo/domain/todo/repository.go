package todo

import "github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/identifier"

type Repository interface {
	Save(t *Todo) error
	Find(i identifier.Identifier) (*Todo, error)
	Delete(i identifier.Identifier) error
}
