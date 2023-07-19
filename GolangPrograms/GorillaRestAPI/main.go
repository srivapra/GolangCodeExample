package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	Name        string `json:"Name"`
	Designation string `json:"Designation"`
	EmpId       string `json:"EmpId"`
}

type Employees []Employee

var employees = Employees{
	{
		Name:        "Prashant",
		Designation: "Software Engineer",
		EmpId:       "12345",
	},
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var newEmp Employee
	dataBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error during reading the request : ", err)
	}

	json.Unmarshal(dataBytes, &newEmp)
	employees = append(employees, newEmp)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmp)

}

func GetEmpBasedOnId(w http.ResponseWriter, r *http.Request) {

	empId := mux.Vars(r)["id"]

	for _, empData := range employees {
		if empData.EmpId == empId {
			json.NewEncoder(w).Encode(empData)

		}
	}
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func UpdateEmployeeBasedOnId(w http.ResponseWriter, r *http.Request) {
	empId := mux.Vars(r)["id"]
	var emp Employee
	dataByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error during reading the request : ", err)
	}

	json.Unmarshal(dataByte, &emp)

	for i, empData := range employees {
		if empData.EmpId == empId {
			empData.Designation = emp.Designation
			employees = append(employees[:i], empData)
			json.NewEncoder(w).Encode(empData)
		}
	}

}

func DeleteEmpDataBasedOnId(w http.ResponseWriter, r *http.Request) {

	empId := mux.Vars(r)["id"]
	for i, empData := range employees {
		if empData.EmpId == empId {
			employees = append(employees[:i], employees[i+1:]...)

		}
	}

}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page!!")
}

// In this case, we have created a function that will return the Welcome to Home Page!!.
// This is when the “/” endpoint is hit.
func main() {

	router := mux.NewRouter().SkipClean(true)
	router.HandleFunc("/", HomePage)
	router.HandleFunc("/emp", CreateEmployee).Methods("POST")
	router.HandleFunc("/allemp", GetAllEmployee).Methods("GET")
	router.HandleFunc("/emp/{id}", GetEmpBasedOnId).Methods("GET")
	router.HandleFunc("/emp/{id}", UpdateEmployeeBasedOnId).Methods("PUT")
	router.HandleFunc("/emp/{id}", DeleteEmpDataBasedOnId).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
