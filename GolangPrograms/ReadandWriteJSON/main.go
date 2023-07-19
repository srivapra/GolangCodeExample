package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

// writeJSONToFile writes a struct to a JSON file
func writeJSONToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

// readJSONFromFile reads a JSON file into a struct
func readJSONFromFile(filename string) (Person, error) {
	var person Person

	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return person, err
	}

	err = json.Unmarshal(jsonFile, &person)
	if err != nil {
		return person, err
	}

	return person, nil
}

func main() {
	// Create a Person struct
	person := Person{Name: "John", Age: 30, Gender: "Male"}

	// Write to JSON file
	err := writeJSONToFile("person.json", person)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}
	fmt.Println("JSON file written successfully")

	// Read from JSON file
	personFromFile, err := readJSONFromFile("person.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	fmt.Println("Person from JSON file:", personFromFile)
}
