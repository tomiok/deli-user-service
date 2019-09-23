package main

import (
	"deli/user-service/engine"
	"github.com/go-chi/chi"
	"net/http"
)

func routes(e engine.Spec, router *chi.Mux) {
	router.Route("/users", func(r chi.Router) {
		r.Get("/{userId}", getUSerByIdHandler)
		r.Post("/admin", func(w http.ResponseWriter, r *http.Request) {
			createsAdminHandler(e, w, r)
		})
		r.Post("/writer", createsWriterHandler)
	})
}
