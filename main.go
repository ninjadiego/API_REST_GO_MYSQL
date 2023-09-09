package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct{
	ID string `json:"id,omitemty"`
	FirstName string `json:"id,omitemty"`
	ID string `json:"id,omitemty"`
	ID string `json:"id,omitemty"`
}

    func GetpeopleEndpoint(w http.ResponseWriter, req *http.Request) {

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
	
	// endpoints
	router.HandleFunc("/people", GetpeopleEndpoint) .Methods("GET")
	router.HandleFunc("/people/{id}", GetPersoneEndpoint) .Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint) .Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint) .Methods("DELETE")

	log.Fatal(http.listenAndServer(":3000", router))

 }