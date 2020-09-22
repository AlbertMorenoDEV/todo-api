package main

import (
	"flag"
	"fmt"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/infrastructure/route"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/infrastructure/server"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/create"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/find"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/infrastructure/persistence/inmemory"
	uiCreate "github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/ui/create"
	uiFind "github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/ui/find"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/command"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/query"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var (
		defaultHost    = os.Getenv("TODOAPI_SERVER_HOST")
		defaultPort, _ = strconv.Atoi(os.Getenv("TODOAPI_SERVER_PORT"))
	)

	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	flag.Parse()

	httpAddr := fmt.Sprintf("%s:%d", *host, *port)

	todoRep := inmemory.NewRepository(todo.Todos{})
	todoCreator := create.NewService(todoRep)
	todoFinder := find.NewService(todoRep)

	cb := command.NewBus()
	if err := cb.RegisterHandler(create.Command{}, create.NewCommandHandler(todoCreator)); err != nil {
		log.Fatal(err)
	}

	qb := query.NewBus()
	if err := qb.RegisterHandler(find.Query{}, find.NewQueryHandler(todoFinder)); err != nil {
		log.Fatal(err)
	}

	s := server.New()
	if err := s.AddRoute(route.New("/todos", "POST", "TodoCreate", uiCreate.NewHandler(cb).Handler)); err != nil {
		log.Fatal(err)
	}
	if err := s.AddRoute(route.New("/todos/{todoId}", "GET", "TodoFind", uiFind.NewHandler(qb).Handler)); err != nil {
		log.Fatal(err)
	}

	fmt.Println("The gopher server is on tap now:", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, s.Router()))
}
