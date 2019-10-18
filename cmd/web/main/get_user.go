package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func getUSerByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "userId")
	u := spec.GetById(id)

	//u == nil means that no records in the database
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
