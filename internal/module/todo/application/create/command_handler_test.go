package create_test

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/create"
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

func TestCreateSuccess(t *testing.T) {
	tests := []struct {
		id        string
		title     string
		due       time.Time
		completed bool
	}{
		{"57b4e893-a946-4c65-baa9-e1653585f731", "In one hour", time.Now().Add(time.Hour), false},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			command := create.Command{
				ID:    test.id,
				Title: test.title,
				Due:   test.due,
			}
			todos := todo.Todos{}
			repository := inmemory.NewRepository(todos)
			service := create.NewService(repository)
			handler := create.NewCommandHandler(service)

			err := handler.Handle(command)
			assert.NoError(t, err)

			todoShouldExist(t, repository, test.id, test.title, test.due, test.completed)
		})
	}
}

func TestCreateFail(t *testing.T) {
	tests := []struct {
		id        string
		title     string
		due       time.Time
		completed bool
	}{
		{"1d6d66f3-1dff-4d29-8bff-34dad612b71b", "One hour ago", time.Now().Add(-1 * time.Hour), false},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			command := create.Command{
				ID:    test.id,
				Title: test.title,
				Due:   test.due,
			}
			todos := todo.Todos{}
			repository := inmemory.NewRepository(todos)
			service := create.NewService(repository)
			handler := create.NewCommandHandler(service)

			err := handler.Handle(command)
			assert.Error(t, err)

			todoShouldNotExist(t, repository, test.id)
		})
	}
}

func todoShouldExist(t *testing.T, repository todo.Repository, idRaw string, titRaw string, duRaw time.Time, coRaw bool) {
	id, err := identifier.New(idRaw)
	assert.NoError(t, err)

	tit, err := title.New(titRaw)
	assert.NoError(t, err)

	du, err := due.New(duRaw)
	assert.NoError(t, err)

	co, err := completed.New(coRaw)
	assert.NoError(t, err)

	found, err := repository.Find(id)
	assert.NoError(t, err)

	assert.True(t, found.ID().EqualsTo(id), "Wrong ResponseID")
	assert.True(t, found.Title().EqualsTo(tit), "Wrong title")
	assert.True(t, found.Due().EqualsTo(du), "Wrong due time")
	assert.True(t, found.Completed().EqualsTo(co), "Wrong completed value")
}

func todoShouldNotExist(t *testing.T, repository todo.Repository, idRaw string) {
	id, err := identifier.New(idRaw)
	assert.NoError(t, err)

	found, err := repository.Find(id)
	assert.Nil(t, found)
	assert.Error(t, err)
}
