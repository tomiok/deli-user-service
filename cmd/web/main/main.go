package main

import (
	"fmt"
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
	fixed  = "8030"
	dbPath = "%s:%s@tcp(%s:3306)/deli_users?parseTime=true"
	)

func main() {
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbURL := os.Getenv("DB_URL")
	if dbPass == "" && dbUser == ""{
		dbPass,dbUser, dbURL = "root", "root", "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = fixed
	}
	logs.InitDefault()
	logs.Infof("CPUs: %d", runtime.NumCPU())

	mux := chi.NewRouter()
	connection := createConnection(fmt.Sprintf(dbPath, dbUser, dbPass, dbURL))

	saveRepo := &datastore.SaveUserRepo{
			DS: connection,
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
	logs.Infof("server running in port %s", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		panic("cannot initialize the server due to: " + err.Error())
	}
}

func createConnection(conn string) *datastore.MysqlDS {
	DS, err := datastore.NewMysqlDS(conn)

	if err != nil {
		panic(err)
	}

	return DS
}
