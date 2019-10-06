package main

import (
	"github.com/deli/user-service/engine"
	"github.com/go-chi/chi"
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
		r.Post("/writer", createsWriterHandler)
	})
}
