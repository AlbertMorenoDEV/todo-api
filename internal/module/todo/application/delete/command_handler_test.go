package delete_test

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/delete"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/infrastructure/persistence/inmemory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDeleteSuccess(t *testing.T) {
	tests := []struct {
		id        string
		title     string
		due       int64
		completed bool
	}{
		{"57b4e893-a946-4c65-baa9-e1653585f731", "In one hour", time.Now().Add(time.Hour).Unix(), false},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			command := delete.Command{ID: test.id}
			todos := todo.Todos{}
			repository := inmemory.NewRepository(todos)
			expTod, err := todo.LoadTodo(test.id, test.title, test.due, test.completed)
			assert.NoError(t, err)
			assert.NoError(t, repository.Save(expTod))
			service := delete.NewService(repository)
			handler := delete.NewCommandHandler(service)

			assert.NoError(t, handler.Handle(command))

			todoShouldNotExist(t, repository, test.id)
		})
	}
}

func TestDeleteFail(t *testing.T) {
	tests := []struct {
		id        string
		title     string
		due       int64
		completed bool
	}{
		{"1d6d66f3-1dff-4d29-8bff-34dad612b71b", "One hour ago", time.Now().Add(-1 * time.Hour).Unix(), false},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			command := delete.Command{ID: test.id}
			todos := todo.Todos{}
			repository := inmemory.NewRepository(todos)
			service := delete.NewService(repository)
			handler := delete.NewCommandHandler(service)

			assert.Error(t, handler.Handle(command))
		})
	}
}

func todoShouldNotExist(t *testing.T, repository todo.Repository, idRaw string) {
	id, err := identifier.New(idRaw)
	assert.NoError(t, err)

	found, err := repository.Find(id)
	assert.Nil(t, found)
	assert.Error(t, err)
}
