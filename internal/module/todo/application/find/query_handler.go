package find

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/identifier"
)

type QueryHandler struct {
	service Service
}

func NewQueryHandler(service Service) *QueryHandler {
	return &QueryHandler{service}
}

func (h *QueryHandler) Handle(query Query) (*Response, error) {
	id, err := identifier.New(query.ID)
	if err != nil {
		return nil, err
	}

	todo, err := h.service.FindById(id)
	if err != nil {
		return nil, err
	}

	return NewResponse(*todo), nil
}
