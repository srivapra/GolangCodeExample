package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Person struct {
	Name   string `yaml:"name"`
	Age    int    `yaml:"age"`
	Gender string `yaml:"gender"`
}

func main() {
	// Create a Person struct
	person := Person{Name: "John", Age: 30, Gender: "Male"}

	// Write to YAML file
	err := writeYAMLToFile("person.yaml", person)
	if err != nil {
		fmt.Println("Error writing YAML file:", err)
		return
	}
	fmt.Println("YAML file written successfully")

	// Read from YAML file
	personFromFile, err := readYAMLFromFile("person.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}
	fmt.Println("Person from YAML file:", personFromFile)
}

// writeYAMLToFile writes a struct to a YAML file
func writeYAMLToFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(yamlData)
	if err != nil {
		return err
	}
	return nil
}

// readYAMLFromFile reads a YAML file into a struct
func readYAMLFromFile(filename string) (Person, error) {
	var person Person

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return person, err
	}

	err = yaml.Unmarshal(yamlFile, &person)
	if err != nil {
		return person, err
	}

	return person, nil
}
