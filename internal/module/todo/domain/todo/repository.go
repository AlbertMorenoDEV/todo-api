package todo

import "github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"

type Repository interface {
	Save(t *Todo) error
	Find(i identifier.Identifier) (*Todo, error)
}
