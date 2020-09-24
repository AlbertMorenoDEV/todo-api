package inmemory

import (
	"fmt"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
	"sync"
)

type Repository struct {
	mtx   sync.RWMutex
	todos todo.Todos
}

func NewRepository(todos todo.Todos) *Repository {
	if todos == nil {
		todos = make(todo.Todos)
	}

	return &Repository{todos: todos}
}

func (r *Repository) Save(t *todo.Todo) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.todos[t.ID().String()] = *t
	return nil
}

func (r *Repository) Find(i identifier.Identifier) (*todo.Todo, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.todos {
		if v.ID().EqualsTo(i) {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("todo item with ID %s could not be found", i.String())
}

func (r *Repository) Delete(i identifier.Identifier) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	_, ok := r.todos[i.String()]
	if ok {
		delete(r.todos, i.String())
		return nil
	}

	return fmt.Errorf("todo item with ID %s could not be found", i.String())
}
