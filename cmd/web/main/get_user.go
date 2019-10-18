package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func getUSerByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	u := spec.GetById(id)
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
