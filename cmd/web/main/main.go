package main

import (
	"github.com/deli/user-service/commons/logs"
	"github.com/deli/user-service/datastore"
	"github.com/deli/user-service/engine"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

const (
	port   = ":8080"
	dbPath = "root:root@tcp(localhost:3306)/deli_user?parseTime=true"
)

func main() {
	port := os.Getenv("PORT")
	logs.InitDefault()
	logs.Infof("CPUs: %d", runtime.NumCPU())

	mux := chi.NewRouter()
	//connection := createConnection(dbPath)

	saveRepo := &datastore.SaveUserRepo{
	//	DS: connection,
	}
	e := engine.New(saveRepo)

	Routes(e, mux)
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
		panic("cannot initialize the server" + err.Error())
	}
}

func createConnection(conn string) *datastore.MysqlDS {
	DS, err := datastore.NewMysqlDS(conn)

	if err != nil {
		panic(err)
	}

	return DS
}
