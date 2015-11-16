package main

import (
	"encoding/json"
	"fmt"
	"github.com/dhanush/findfalcone/Godeps/_workspace/src/github.com/gorilla/mux"
	"net/http"
	"os"
)

var planets = Planets{Planet{"Donlon", 400}, Planet{"Enchai", 40}, Planet{"Jebing", 100}, Planet{"Sapir", 240}, Planet{"Lerbin", 200}, Planet{"Pingasor", 80}}

var vehicles = Vehicles{Vehicle{"Spaceship", 5, 100, 50}, Vehicle{"Rocket", 3, 200, 100}, Vehicle{"Cycle", 8, 40, 10}, Vehicle{"Missile", 1, 300, 150}}

func Init(rw http.ResponseWriter, req *http.Request) {

}

func PlanetsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(planets); err != nil {
		panic(err)
	}
}

func VehicleHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(vehicles); err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Starting server on " + port)
	r.HandleFunc("/init", Init).Methods("POST").Headers("Accept", "application/json")
	r.HandleFunc("/planets", PlanetsHandler).Methods("GET")
	r.HandleFunc("/vehicles", VehicleHandler).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
