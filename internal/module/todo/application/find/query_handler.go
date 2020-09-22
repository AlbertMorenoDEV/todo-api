package find

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
)

type QueryHandler struct {
	service Service
}

func NewQueryHandler(s Service) *QueryHandler {
	return &QueryHandler{s}
}

func (h *QueryHandler) Handle(q Query) (*Response, error) {
	id, err := identifier.New(q.ID)
	if err != nil {
		return nil, err
	}

	todo, err := h.service.FindById(id)
	if err != nil {
		return nil, err
	}

	return NewResponse(*todo), nil
}
