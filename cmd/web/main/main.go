package main

import (
	"fmt"
	"github.com/deli/user-service/commons/logs"
	"github.com/deli/user-service/datastore"
	"github.com/deli/user-service/engine"
	"github.com/deli/user-service/token"
	"github.com/go-chi/chi"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

const (
	fixed  = "8080"
	dbPath = "%s:%s@tcp(%s:3306)/deli_users?parseTime=true"
	secret = "isasecret"
)

var spec engine.Spec

func main() {
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbURL := os.Getenv("DB_URL")
	if dbPass == "" && dbUser == "" {
		dbPass, dbUser, dbURL = "root", "root", "localhost"
	}
	// init token
	initJWT(secret)
	port := os.Getenv("PORT")
	if port == "" {
		port = fixed
	}
	logs.InitDefault()
	logs.Infof("CPUs: %d", runtime.NumCPU())

	mux := chi.NewRouter()
	connection := createConnection(fmt.Sprintf(dbPath, dbUser, dbPass, dbURL))

	userRepository := &datastore.UserRepository{
		DS: connection,
	}
	 spec = engine.New(userRepository)

	Routes(mux)
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

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err := s.ListenAndServe()
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

func initJWT(secret string) {
	token.Init(secret)
}
