package main

import (
	"encoding/json"
	"fmt"
	"github.com/deli/user-service/commons"
	"github.com/deli/user-service/commons/logs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/docgen/raml"
	"github.com/go-chi/render"
	"net/http"
	"runtime"
	"time"
)

func Routes(router *chi.Mux) {

	router.Use(render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.Recoverer,
		middleware.RequestID,
		middleware.DefaultCompress,
		middleware.Heartbeat("/ping"))

	router.Route("/users", func(rt chi.Router) {

		rt.Get("/{userId}", getUSerByIdHandler)
		rt.Post("/aw", createsAdminOrWriterHandler)
		rt.Post("/login", LoginHandler)
	})

	router.Get("/health", healthCheck)

	if false {
		fmt.Println(generateRaml(router))
	}
}

//LoginHandler.
//Returns a JWT token.
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	validateUserHandler(w, r, encrypt)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"status":      "OK",
		"server_time": time.Now(),
		"server_OS":   runtime.GOOS,
		"server_arch": runtime.GOARCH,
		"CPUs":        runtime.NumCPU(),
	})
}

func encrypt(r *http.Request) (string, string) {
	type credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	cred := credentials{}
	_ = json.NewDecoder(r.Body).Decode(&cred)

	return cred.Username, commons.EncryptPass(cred.Password)
}

func customLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.Info("Executing middlewareOne")
		next.ServeHTTP(w, r)
		logs.Info("Executing middlewareOne again")
	})
}

func contentTypeM(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func generateRaml(r chi.Routes) *raml.RAML {

	ramlDocs := &raml.RAML{
		Title:     "Deli User Service",
		BaseUri:   "https://xxxxx.herokuapp.com/users",
		Version:   "v1.0",
		MediaType: "application/json",
	}

	_ = chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {

		funcInfo := docgen.GetFuncInfo(handler)
		resource := &raml.Resource{
			DisplayName: funcInfo.Func,
			Description: funcInfo.Comment,
			//Responses:       map[int]raml.Response{200: raml.Response{}},
			//Body:            map[string]raml.Example{"" :raml.Example{Example:""}},
			//Is:              nil,
			//	Type:            "",
			SecuredBy:       nil,
			UriParameters:   nil,
			QueryParameters: nil,
			Resources:       nil,
		}

		return ramlDocs.Add(method, route, resource)
	})

	return ramlDocs
}
