package main

import (
	"encoding/json"
	"fmt"
	"github.com/deli/user-service/engine"
	"github.com/go-chi/chi"
	"github.com/go-chi/docgen"
	"net/http"
)

func Routes(e engine.Spec, router *chi.Mux) {
	router.Route("/users", func(r chi.Router) {
		r.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
			getUSerByIdHandler(e, w, r)
		})
		r.Post("/aw", func(w http.ResponseWriter, r *http.Request) {
			createsAdminOrWriterHandler(e, w, r)
		})

		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {

		})
	})

	router.Get("/", healthCheck)

	if true {
		doc := docgen.JSONRoutesDoc(router)
		fmt.Println(doc)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"status": "OK"})
}
