package create

import (
	"errors"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/due"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/title"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/command"
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

	title, err := title.New(cmd.Title)
	if err != nil {
		return err
	}

	due, err := due.New(cmd.Due)
	if err != nil {
		return err
	}

	return h.service.CreateTodo(id, title, due)
}
