package main

import (
	"encoding/json"
	"github.com/deli/user-service/commons/logs"
	"github.com/deli/user-service/models"
	"io/ioutil"
	"net/http"
)

func createsAdminOrWriterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	userType := r.URL.Query()["type"][0]

	defer r.Body.Close()

	id, err := spec.Save(userFunction(r, userType))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := make(map[string]string)
	res["user_id"] = id
	_ = json.NewEncoder(w).Encode(res)
}

func userFunction(r *http.Request, userType string) func() *models.User {

	//TODO finish this error handling :(
	return func() *models.User {
		user := models.User{}
		user.UserType = &models.UserType{Title: userType}

		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &user)

		logs.Info("Mapping user")
		return user.DoMap()
	}
}
