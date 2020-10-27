package todo

import (
	"fmt"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/completed"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/due"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/title"
)

type Todo struct {
	id        identifier.Identifier
	title     title.Title
	completed completed.Completed
	due       due.Due
}

type Todos map[string]Todo

func NewTodo(id identifier.Identifier, tit title.Title, due due.Due) (*Todo, error) {
	if due.IsPast() {
		return nil, fmt.Errorf("Wrong due value: %d, can not be in the past", due.Time().Unix())
	}

	return &Todo{
		id:        id,
		title:     tit,
		completed: completed.False(),
		due:       due,
	}, nil
}

func LoadTodo(idRaw string, titRaw string, dueRaw int64, comRaw bool) (*Todo, error) {
	id, err := identifier.New(idRaw)
	if err != nil {
		return nil, err
	}

	tit, err := title.New(titRaw)
	if err != nil {
		return nil, err
	}

	d, err := due.FromMilliseconds(dueRaw)
	if err != nil {
		return nil, err
	}

	com, err := completed.New(comRaw)
	if err != nil {
		return nil, err
	}

	return &Todo{
		id:        id,
		title:     tit,
		completed: com,
		due:       *d,
	}, nil
}

func (t *Todo) ID() identifier.Identifier {
	return t.id
}

func (t *Todo) Title() title.Title {
	return t.title
}

func (t *Todo) Completed() completed.Completed {
	return t.completed
}

func (t *Todo) Due() due.Due {
	return t.due
}
