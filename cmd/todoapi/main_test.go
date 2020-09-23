package main

import (
	"fmt"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/ui/create"
	"github.com/gavv/httpexpect/v2"
	"github.com/icrowley/fake"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthCheck(t *testing.T) {
	r := createRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/").
		Expect().
		Status(http.StatusOK)
}

func TestCreateAndGetTodo(t *testing.T) {
	r := createRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	req := create.Request{
		ID:    uuid.NewV4().String(),
		Title: fake.Title(),
		Due:   time.Now().Add(7 * 24 * time.Hour),
	}

	e.POST("/todos").WithJSON(req).
		Expect().
		Status(http.StatusCreated)

	schema := `{
		"type": "object",
		"properties": {
			"id":        	{"type": "string"},
			"title": 		{"type": "string"},
			"completed":	{"type": "boolean"},
			"due": 			{"type": "string"}
		},
		"required": ["id", "title", "completed", "due"]
	}`

	res := e.GET(fmt.Sprintf("/todos/%s", req.ID)).
		Expect().
		Status(http.StatusOK).JSON()

	res.Schema(schema)
	res.Object().Value("id").Equal(req.ID)
	res.Object().Value("title").Equal(req.Title)
	res.Object().Value("due").Equal(req.Due)
	res.Object().Value("completed").Equal(false)
}
