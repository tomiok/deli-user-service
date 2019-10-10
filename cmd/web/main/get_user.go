package main

import (
	"encoding/json"
	"github.com/deli/user-service/engine"
	"github.com/go-chi/chi"
	"net/http"
)

func getUSerByIdHandler(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	u := e.GetById(id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(u)

}
