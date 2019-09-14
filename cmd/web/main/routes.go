package main

import (
	"github.com/go-chi/chi"
)

func routes(router *chi.Mux) {
	router.Route("/users", func(r chi.Router) {
		r.Get("/{userId}", getUSerByIdHandler)
		r.Post("/admin", createsAdminHandler)
		r.Post("/writer", createsWriterHandler)
	})
}
