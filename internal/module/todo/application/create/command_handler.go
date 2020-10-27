package create

import (
	"errors"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/due"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/title"
	"github.com/AlbertMorenoDEV/todo-api/pkg/infrastructure/bus/command"
)

type CommandHandler struct {
	service Service
}

func NewCommandHandler(service Service) *CommandHandler {
	return &CommandHandler{service}
}

func (h *CommandHandler) Handle(c command.Command) error {
	cmd, ok := c.(Command)
	if !ok {
		return errors.New("invalid command")
	}

	id, err := identifier.New(cmd.ID)
	if err != nil {
		return err
	}

	tit, err := title.New(cmd.Title)
	if err != nil {
		return err
	}

	d, err := due.FromMilliseconds(cmd.Due)
	if err != nil {
		return err
	}

	return h.service.CreateTodo(id, tit, *d)
}
