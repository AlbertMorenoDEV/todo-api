package main

import (
	"flag"
	"fmt"
	"github.com/AlbertMorenoDEV/go-ddd-playground/internal/infrastructure/server"
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

	s := server.New()

	// s.AddRoute()

	fmt.Println("The gopher server is on tap now:", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, s.Router()))
}
