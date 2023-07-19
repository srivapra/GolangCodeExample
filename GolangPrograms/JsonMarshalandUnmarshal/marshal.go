// Converting a Go struct in json is calling marshaling in golang and
// converting a json into a go struct is callled unmarshaling in golang

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Address string
}

func main() {
	emp := Employee{Name: "George Smith", Age: 30, Address: "Newyork, USA"}
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(empData))
}
