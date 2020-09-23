package main

import (
	"flag"
	"fmt"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/create"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/application/find"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/domain/todo"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/infrastructure/persistence/inmemory"
	uiCreate "github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/ui/create"
	uiFind "github.com/AlbertMorenoDEV/go-ddd-playground/internal/module/todo/ui/find"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/command"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/bus/query"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/http/server"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/infrastructure/http/server/route"
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/ui/health"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
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

	r := createRouter()

	fmt.Println("Running todoapi server on:", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, r))
}

func createRouter() http.Handler {
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

	if err := s.AddRoute(route.New("/", "GET", "HealthCheck", health.Handler)); err != nil {
		log.Fatal(err)
	}

	createTodoHandler := uiCreate.NewHandler(cb)
	if err := s.AddRoute(route.New("/todos", "POST", "TodoCreate", createTodoHandler.Handler)); err != nil {
		log.Fatal(err)
	}

	findTodoHandler := uiFind.NewHandler(qb)
	if err := s.AddRoute(route.New("/todos/{todoId}", "GET", "TodoFind", findTodoHandler.Handler)); err != nil {
		log.Fatal(err)
	}

	return s.Router()
}
