package main

import (
	"encoding/json"
	"fmt"
	"github.com/dhanush/findfalcone/Godeps/_workspace/src/github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//no of balls faced or bowled in test cricket - used to calculate the
const (
	DRAVID int = 31258
	SACHIN int = 29437
	KALLIS int = 28903
	MURALI int = 44039
	KUMBLE int = 40850
	WARNE  int = 40705
)

var total_balls = []int{DRAVID, SACHIN, KALLIS, MURALI, KUMBLE, WARNE}

var falcones map[string]int = make(map[string]int)

var planets = Planets{Planet{"Donlon", 400}, Planet{"Enchai", 40}, Planet{"Jebing", 100}, Planet{"Sapir", 240}, Planet{"Lerbin", 200}, Planet{"Pingasor", 80}}

var vehicles = Vehicles{Vehicle{"Space pod", 5, 100, 50}, Vehicle{"Space rocket", 3, 200, 100}, Vehicle{"Space shuttle", 8, 40, 10}, Vehicle{"Space ship", 1, 300, 150}}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

///returns a token for an N integer
func randSeq(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//returns a random no out of 0-6
func where_is_falcone() int {
	rand.Seed(time.Now().UTC().UnixNano())
	var no = rand.Intn(6)
	return (total_balls[no] * rand.Intn(10)) % 6
}

//returns a token for a user who is trying to find falcone.
func Init(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusOK)
	var random_str = randSeq(32)
	var falcone_planet = where_is_falcone()
	falcones[random_str] = falcone_planet
	var token = map[string]string{"token": random_str}
	if err := json.NewEncoder(rw).Encode(token); err != nil {
		panic(err)
	}
}

//returns all the planets
func PlanetsHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(planets); err != nil {
		panic(err)
	}
}

//returns all the vehicles
func VehicleHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(vehicles); err != nil {
		panic(err)
	}
}

//API to find the falcone
func FindFalcone(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(req.Body)
	var find_falcone FindFalconeReq
	err := decoder.Decode(&find_falcone)
	if err != nil {
		errorHandler(rw, req, http.StatusBadRequest, err.Error())
		return
	}
	var planetNames = find_falcone.PlanetNames
	if len(planetNames) != 4 {
		errorHandler(rw, req, http.StatusBadRequest, "No of Planet names has to be 4")
		return
	}
	var vehicleNames = find_falcone.VehicleNames
	if len(vehicleNames) != 4 {
		errorHandler(rw, req, http.StatusBadRequest, "No of Vehicle names has to be 4")
		return
	}
	if len(falcones) == 0 {
		errorHandler(rw, req, http.StatusBadRequest, "Token not initialized. Please get a new token with the /token API")
		return
	}

	// var falconePlanetIndex = falcones[find_falcone.Token]
	if falconePlanetIndex, ok := falcones[find_falcone.Token]; ok {
		rw.WriteHeader(http.StatusOK)
		var falconePlanet = planets[falconePlanetIndex]
		for _, name := range planetNames {
			if name == falconePlanet.Name {
				var status = map[string]string{"status": "success"}
				if err := json.NewEncoder(rw).Encode(status); err != nil {
					panic(err)
				}
				return
			}
		}
	} else {
		errorHandler(rw, req, http.StatusBadRequest, "Token not initialized. Please get a new token with the /token API")
		return
	}
	var status = map[string]string{"status": "false"}
	if err := json.NewEncoder(rw).Encode(status); err != nil {
		panic(err)
	}
}

func errorHandler(rw http.ResponseWriter, req *http.Request, status int, message string) {
	rw.WriteHeader(status)
	var error = map[string]string{"error": message}
	if err := json.NewEncoder(rw).Encode(error); err != nil {
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
	r.HandleFunc("/token", Init).Methods("POST").Headers("Accept", "application/json")
	r.HandleFunc("/planets", PlanetsHandler).Methods("GET")
	r.HandleFunc("/vehicles", VehicleHandler).Methods("GET")
	r.HandleFunc("/find", FindFalcone).Methods("POST").Headers("Accept", "application/json")

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
