package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var planets = Planets{Planet{"Donlon", 400}, Planet{"Enchai", 40}, Planet{"Jebing", 100}, Planet{"Sapir", 240}, Planet{"Lerbin", 200}, Planet{"Pingasor", 80}}

func Init(rw http.ResponseWriter, req *http.Request) {

}

func PlanetsHandler(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(planets); err != nil {
		panic(err)
	}

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/init", Init).Methods("POST").Headers("Accept", "application/json")
	r.HandleFunc("/planets", PlanetsHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
