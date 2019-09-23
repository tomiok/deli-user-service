package main

import (
	"deli/user-service/engine"
	"deli/user-service/model"
	"github.com/labstack/gommon/log"
	"net/http"
)

func createsAdminHandler(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	u := &model.User{
		Name:         "Tomas",
		LastName:     "Lingotti",
		EmailAddress: "tomi@msn.com",
		UserType:     &model.UserType{Title: "admin"},
		Password:     "epicpass",
		City:         "Rosario",
		Country:      "ARG",
	}

	log.Info("Mapping user")
	e.Save(u.DoMap())
}

func createsWriterHandler(writer http.ResponseWriter, request *http.Request) {

}
