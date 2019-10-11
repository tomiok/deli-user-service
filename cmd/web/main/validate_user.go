package main

import (
	"encoding/json"
	"github.com/deli/user-service/engine"
	"net/http"
)

func validateUserHandler(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	token, err := e.ValidateUser("", "")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(responseBadRequest(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tokenResponse(token))
}

func tokenResponse(s string) map[string]string {
	return map[string]string{"JWT": s}
}

func responseBadRequest(s string) map[string]string {
	return map[string]string{
		"operation_status": "FAILED",
		"reason":           s,
	}
}
