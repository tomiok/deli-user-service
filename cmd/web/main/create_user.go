package main

import (
	"deli/user-service/engine"
	"deli/user-service/model"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"net/http"
)

func createsAdminHandler(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	u := model.User{}
	u.UserType = &model.UserType{Title: "admin"}

	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &u)
	log.Info("Mapping user")

	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	e.Save(u.DoMap())
}

func createsWriterHandler(writer http.ResponseWriter, request *http.Request) {

}
