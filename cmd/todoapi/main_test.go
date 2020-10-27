package main

import (
	"fmt"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/todo"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/infrastructure/persistence/inmemory"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/ui/create"
	"github.com/gavv/httpexpect/v2"
	"github.com/icrowley/fake"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthCheck(t *testing.T) {
	todoRep := inmemory.NewRepository(todo.Todos{})
	r := createRouter(todoRep)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/").
		Expect().
		Status(http.StatusOK)
}

func TestCreateTodo(t *testing.T) {
	todoRep := inmemory.NewRepository(todo.Todos{})
	r := createRouter(todoRep)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	expTodo := randomTodo(t)
	req := create.Request{
		ID:    expTodo.ID().String(),
		Title: expTodo.Title().String(),
		Due:   expTodo.Due().Time().Unix(),
	}

	e.POST("/todos").WithJSON(req).
		Expect().
		Status(http.StatusCreated)

	foundTodo, err := todoRep.Find(expTodo.ID())
	assert.NotNil(t, foundTodo)
	assert.NoError(t, err)
	assert.Equal(t, &expTodo, foundTodo)
}

func TestGetTodo(t *testing.T) {
	todoRep := inmemory.NewRepository(todo.Todos{})
	r := createRouter(todoRep)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	expTodo := randomTodo(t)
	assert.NoError(t, todoRep.Save(&expTodo))

	schema := `{
		"type": "object",
		"properties": {
			"id":        	{"type": "string"},
			"title": 		{"type": "string"},
			"completed":	{"type": "boolean"},
			"due": 			{"type": "integer"}
		},
		"required": ["id", "title", "completed", "due"]
	}`

	res := e.GET(fmt.Sprintf("/todos/%s", expTodo.ID())).
		Expect().
		Status(http.StatusOK).JSON()

	res.Schema(schema)
	res.Object().Value("id").Equal(expTodo.ID().String())
	res.Object().Value("title").Equal(expTodo.Title().String())
	res.Object().Value("due").Equal(expTodo.Due().Time().Unix())
	res.Object().Value("completed").Equal(expTodo.Completed().Bool())
}

func TestDeleteTodo(t *testing.T) {
	todoRep := inmemory.NewRepository(todo.Todos{})
	r := createRouter(todoRep)

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	tod := randomTodo(t)

	err := todoRep.Save(&tod)
	assert.NoError(t, err)

	e.DELETE(fmt.Sprintf("/todos/%s", tod.ID())).
		Expect().
		Status(http.StatusOK)

	f, err := todoRep.Find(tod.ID())
	assert.Error(t, err)
	assert.Nil(t, f)
}

func randomTodo(t *testing.T) todo.Todo {
	tod, err := todo.LoadTodo(uuid.NewV4().String(), fake.Title(), time.Now().Add(7*24*time.Hour).Unix(), false)
	assert.NoError(t, err)

	return *tod
}
