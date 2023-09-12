package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct{
	ID string `json:"id,omitemty"`
	FirstName string `json:"firstname,omitemty"`
	LastName string `json:"lastname,omitemty"`
	Address *Address `json:"address,omitemty"`
}

type Andress struct{
	City string `json:"City,omitempy"`
	State string `json:"State,omitempy"`
}

var people []person

    func GetpeopleEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(people)
	}
	
	func GetPersoneEndpoint(w http.ResponseWriter, req *http.Request) {

	}

	func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

	}

	func DeletePersonEndpointPersonEndpoint(w http.ResponseWriter, req *http.Request) {

	}

	router.HandleFunc("/people", GetpeopleEndpoint) .Methods("GET")

    func main()  {
	router := mux.NewRouter()

	people = append(people, person{ID: "1", FirstName: "Ryan", lastname: "Ray", Address:
	&Address{City: "Dubling", State: "California"}})
	people = append(people, person{ID: "2", FirstName: "joe", lastname: "McMillan"})
	
	// endpoints
	router.HandleFunc("/people", GetpeopleEndpoint) .Methods("GET")
	router.HandleFunc("/people/{id}", GetPersoneEndpoint) .Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint) .Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint) .Methods("DELETE")

	log.Fatal(http.listenAndServer(":3000", router))

 }