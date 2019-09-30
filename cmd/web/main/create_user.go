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
	id, err := e.Save(u.DoMap())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := make(map[string]string)
	res["user_id"] = id
	_ = json.NewEncoder(w).Encode(res)
}

func createsWriterHandler(w http.ResponseWriter, r *http.Request) {

}
