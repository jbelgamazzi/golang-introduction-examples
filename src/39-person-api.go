package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	fmt.Println("servico iniciado")

	people = append(people, Person{ID: "1", Firstname: "jose", Lastname: "sila", Address: &Address{City: "Maringa", State: "PR"}})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}

func GetPeopleEndPoint(w http.ResponseWriter, request *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndPoint(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndPoint(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	var person Person
	_ = json.NewDecoder(request.Body).Decode(&person)

	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePersonEndPoint(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}
