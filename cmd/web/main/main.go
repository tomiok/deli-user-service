package main

import (
	"deli/user-service/datastore"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"os/signal"
)

const (
	port    = ":8080"
	db_path = "root:rootd@tcp(127.0.0.1:3306)/deli_user"
)

func main() {
	mux := chi.NewRouter()
	routes(mux)

	createConnection(db_path)

	go func() {
		startServer(mux, port)
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

func createConnection(conn string) *datastore.MysqlDS {
	DS, err := datastore.NewMysqlDS(conn)

	if err != nil {
		panic(err)
	}

	return DS
}
