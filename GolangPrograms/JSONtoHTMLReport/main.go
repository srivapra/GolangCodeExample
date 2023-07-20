package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

type Person struct {
	Name     string  `json:"name"`
	Language string  `json:"language"`
	ID       string  `json:"id"`
	Bio      string  `json:"bio"`
	Version  float64 `json:"version"`
}

func main() {
	// Read the JSON file
	file, err := os.Open("abc.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse the JSON content
	var person Person
	err = json.Unmarshal(content, &person)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Generate HTML using template
	htmlTemplate := `<html>
	<head>
		<title>HTML Test Report</title>
	</head>
	<body>
		<h1>{{.Name}}</h1>
		<p>Language: {{.Language}}</p>
		<p>ID: {{.ID}}</p>
		<p>Bio: {{.Bio}}</p>
		<p>Version: {{.Version}}</p>
	</body>
</html>`

	tmpl, err := template.New("profile").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Error parsing HTML template:", err)
		return
	}

	// Create output file
	outputFile, err := os.Create("output.html")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Generate HTML using the data and template
	err = tmpl.Execute(outputFile, person)
	if err != nil {
		fmt.Println("Error generating HTML:", err)
		return
	}

	fmt.Println("HTML document generated successfully.")
}
