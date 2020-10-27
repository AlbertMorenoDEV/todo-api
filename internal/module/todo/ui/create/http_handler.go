package create

import (
	"encoding/json"
	"github.com/AlbertMorenoDEV/todo-api/internal/module/todo/application/create"
	"github.com/AlbertMorenoDEV/todo-api/pkg/infrastructure/bus/command"
	"github.com/AlbertMorenoDEV/todo-api/pkg/infrastructure/http/jsonapi"
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
		resp := append(jsonapi.Errors{}, jsonapi.Error{Status: http.StatusInternalServerError, Title: "Error unmarshalling request body"})
		_ = json.NewEncoder(w).Encode(resp)
		log.Print(err)
		return
	}

	cmd := create.Command{ID: req.ID, Title: req.Title, Due: req.Due}

	if err := h.commandBus.Dispatch(cmd); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := append(jsonapi.Errors{}, jsonapi.Error{Status: http.StatusInternalServerError, Title: "Can't create a todo"})
		_ = json.NewEncoder(w).Encode(resp)
		log.Print(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
