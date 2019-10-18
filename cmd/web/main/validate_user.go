package main

import (
	"encoding/json"
	"net/http"
)

type encryptFn func(r *http.Request) (string, string)

func validateUserHandler(w http.ResponseWriter, r *http.Request, fn encryptFn) {
	username, pass := fn(r)
	token, err := spec.ValidateUser(username, pass)

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
