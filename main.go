package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Planet struct {
	Name string `json:"name"`
}

var planets [10]Planet

func Init(rw http.ResponseWriter, req *http.Request) {

}

func Planets(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(planets); err != nil {
		panic(err)
	}

}

func main() {
	fmt.Println("Hello")

	r := mux.NewRouter()
	r.HandleFunc("/init", Init).Methods("POST").Headers("Accept", "application/json")
	r.HandleFunc("/planets", Planets).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
