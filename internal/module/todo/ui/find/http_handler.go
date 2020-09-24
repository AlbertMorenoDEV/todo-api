package find

import (
	"encoding/json"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/find"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/query"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/http/jsonapi"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	queryBus query.Bus
}

func NewHandler(qb query.Bus) Handler {
	return Handler{queryBus: qb}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	qr := find.Query{ID: vars["todoId"]}

	resp, err := h.queryBus.Ask(qr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := append(jsonapi.Errors{}, jsonapi.Error{Status: http.StatusInternalServerError, Title: "Can't find the todo"})
		_ = json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := append(jsonapi.Errors{}, jsonapi.Error{Status: http.StatusInternalServerError, Title: "Internal Server Error"})
		_ = json.NewEncoder(w).Encode(resp)
	}
}
