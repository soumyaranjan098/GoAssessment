package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee", GetEmployees).Methods("GET")
	r.HandleFunc("/employee/{id}", UpdateByID).Methods("PUT")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp employee
	json.NewDecoder(r.Body).Decode(&emp)
	db.Create(&emp)
	json.NewEncoder(w).Encode("New employee created successfully...")
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var studs []employee
	res := db.Find(&studs)
	if res.RowsAffected != 0 {
		json.NewEncoder(w).Encode(studs)
	} else {
		json.NewEncoder(w).Encode("No student found..")
	}

}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee employee
	res := db.First(&employee, mux.Vars(r)["id"])
	if res.RowsAffected != 0 {
		json.NewDecoder(r.Body).Decode(&employee)
		db.Save(employee.Balance)
		json.NewEncoder(w).Encode(employee)
	} else {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("No employee found with the given id")

	}

}
