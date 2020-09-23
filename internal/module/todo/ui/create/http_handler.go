package create

import (
	"encoding/json"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/create"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/command"
	"log"
	"net/http"
)

type Handler struct {
	commandBus command.Bus
}

func NewHandler(commandBus command.Bus) Handler {
	return Handler{commandBus: commandBus}
}

type Request struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Due   int64  `json:"due"`
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error unmarshalling request body")
		log.Print(err)
		return
	}

	cmd := create.Command{
		ID:    req.ID,
		Title: req.Title,
		Due:   req.Due,
	}

	if err := h.commandBus.Dispatch(cmd); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Can't create a todo")
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
