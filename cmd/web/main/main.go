package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	fmt.Println("start application")
	mux := chi.NewRouter()
	routes(mux)

	go func() {

	}()

	// Wait for terminate signal to shut down server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func startServer(mux *chi.Mux, port string) {
	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic("cannot initialize the server")
	}
}
