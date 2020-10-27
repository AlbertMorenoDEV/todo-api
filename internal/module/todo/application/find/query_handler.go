package find

import (
	"errors"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/domain/identifier"
	"github.com/AlbertMorenoDEV/todo-api/pkg/infrastructure/bus/query"
)

type QueryHandler struct {
	service Service
}

func NewQueryHandler(s Service) *QueryHandler {
	return &QueryHandler{s}
}

func (h *QueryHandler) Handle(q query.Query) (*query.Response, error) {
	qr, ok := q.(Query)
	if !ok {
		return nil, errors.New("invalid query")
	}

	id, err := identifier.New(qr.ID)
	if err != nil {
		return nil, err
	}

	todo, err := h.service.FindById(id)
	if err != nil {
		return nil, err
	}

	r := NewResponse(*todo)

	return &r, nil
}
