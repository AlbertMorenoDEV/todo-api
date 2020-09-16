package find_test

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/find"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/completed"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/due"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/title"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/infrastructure/persistence/inmemory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFindSuccess(t *testing.T) {
	tests := []struct {
		id        string
		title     string
		due       time.Time
		completed bool
	}{
		{
			"57b4e893-a946-4c65-baa9-e1653585f731",
			"In one hour",
			time.Now().Add(time.Hour),
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			existing := todoFromRaw(t, test.id, test.title, test.due, test.completed)
			expected := find.NewResponse(existing)
			query := find.Query{ID: test.id}
			todos := todo.Todos{}
			todos[test.id] = existing
			repository := inmemory.NewRepository(todos)
			service := find.NewService(repository)
			handler := find.NewQueryHandler(service)

			res, err := handler.Handle(query)
			assert.NoError(t, err)
			assert.Equal(t, expected, res)
		})
	}
}

func TestFindFail(t *testing.T) {
	tests := []struct {
		id    string
		title string
	}{
		{
			"57b4e893-a946-4c65-baa9-e1653585f731",
			"In one hour",
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			query := find.Query{ID: test.id}
			todos := todo.Todos{}
			repository := inmemory.NewRepository(todos)
			service := find.NewService(repository)
			handler := find.NewQueryHandler(service)

			res, err := handler.Handle(query)
			assert.Nil(t, res)
			assert.Error(t, err)
		})
	}
}

func todoFromRaw(t *testing.T, idRaw string, titRaw string, duRaw time.Time, coRaw bool) todo.Todo {
	id, err := identifier.New(idRaw)
	assert.NoError(t, err)

	tit, err := title.New(titRaw)
	assert.NoError(t, err)

	du, err := due.New(duRaw)
	assert.NoError(t, err)

	co, err := completed.New(coRaw)
	assert.NoError(t, err)

	tod, err := todo.LoadTodo(id, tit, du, co)
	assert.NoError(t, err)

	return *tod
}
