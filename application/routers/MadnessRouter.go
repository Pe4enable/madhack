package routers

import (
	. "../controllers"
	"github.com/gorilla/mux"
)

func MadnessRouter(path string, router *mux.Router) {

	sr := router.PathPrefix(path).Subrouter()
	sr.HandleFunc("/{employeeId}", GetData).Methods("GET")
	sr.HandleFunc("/{employeeId}", GetData).Methods("POST")
}