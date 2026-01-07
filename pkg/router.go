package pkg

import (
	"github.com/gorilla/mux"
)

var R = mux.NewRouter()

func SetRouter() {
	// Set different endpoints
	R.HandleFunc("/list", GetList).Methods("GET")
	R.HandleFunc("/list/{id}", GetTask).Methods("GET")
	R.HandleFunc("/list", AddTask).Methods("POST")
	R.HandleFunc("/list/{id}", UpdateTask).Methods("PATCH")
	R.HandleFunc("/list/{id}", DeleteTask).Methods("DELETE")
}
