package main

import (
	"encoding/json"
	"github.com/deli/user-service/engine"
	"github.com/deli/user-service/logs"
	"github.com/deli/user-service/model"
	"io/ioutil"
	"net/http"
)

func createsAdminOrWriterHandler(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	userType := r.URL.Query()["type"][0]
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	id, err := e.Save(userFunction(r, userType))

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

func userFunction(r *http.Request, userType string) func() *model.User {

	//TODO finish this error handling :(
	return func() *model.User {
		user := model.User{}
		user.UserType = &model.UserType{Title: userType}

		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &user)

		logs.Info("Mapping user")
		return user.DoMap()
	}
}
