package main

import (
	"deli/user-service/engine"
	"net/http"
)

func createsAdminHandler(e engine.Spec, writer http.ResponseWriter, request *http.Request) {
	e.SaveUser()
}

func createsWriterHandler(writer http.ResponseWriter, request *http.Request) {

}
