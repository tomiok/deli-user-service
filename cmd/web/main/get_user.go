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
	if u == nil {
		w.WriteHeader(http.StatusNotFound)
		res := map[string]string{
			"operation_status": "FAILED",
			"reason":           "user with id " + id + "not found",
		}
		_ = json.NewEncoder(w).Encode(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(u)
}
