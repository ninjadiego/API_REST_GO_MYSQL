package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"City,omitempty"`
	State string `json:"State,omitempty"`
}

var people []person

func GetpeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewDecoder(req.Body).Decode(&item)
			json.NewEncoder(w).Encode(&item)
			return
		}
	}
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person person
	_ = json.NewDecoder(req.Body).Decode(&person)

	person.ID = params["id"]
	person.FirstName = params["firstname"]

	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()

	people = append(people, person{ID: "1", FirstName: "Ryan", LastName: "Ray", Address: &Address{City: "Dubling", State: "California"}})
	people = append(people, person{ID: "2", FirstName: "joe", LastName: "McMillan"})

	// endpoints
	router.HandleFunc("/people", GetpeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

}
