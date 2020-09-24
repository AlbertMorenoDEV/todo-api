package delete

import (
	"encoding/json"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/delete"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/command"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	commandBus command.Bus
}

func NewHandler(commandBus command.Bus) Handler {
	return Handler{commandBus: commandBus}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	com := delete.Command{ID: vars["todoId"]}

	if err := h.commandBus.Dispatch(com); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Can't delete the todo")
		return
	}

	w.WriteHeader(http.StatusOK)
}
