package create

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/due"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/title"
)

type CommandHandler struct {
	service Service
}

func NewCommandHandler(service Service) *CommandHandler {
	return &CommandHandler{service}
}

func (h *CommandHandler) Handle(command Command) error {
	id, err := identifier.New(command.ID)
	if err != nil {
		return err
	}

	title, err := title.New(command.Title)
	if err != nil {
		return err
	}

	due, err := due.New(command.Due)
	if err != nil {
		return err
	}

	return h.service.CreateTodo(id, title, due)
}
