package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"
)

type Data struct {
	Title string
	Data  map[string]interface{}
}

func main() {
	// Sample JSON data
	jsonData := `{
		"Title": "JSON to HTML Report",
		"Data": {
			"Name": "John Doe",
			"Age": 30,
			"Email": "john.doe@example.com",
			"Address": "123 Main St"
		}
	}`

	// Parse JSON data into the Data struct
	var data Data
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		log.Fatal("Failed to parse JSON:", err)
	}

	// Load the HTML template from file
	tmpl, err := template.ParseFiles("report_template.html")
	if err != nil {
		log.Fatal("Failed to parse HTML template:", err)
	}

	// Create the output HTML file
	outputFile, err := os.Create("report.html")
	if err != nil {
		log.Fatal("Failed to create output file:", err)
	}
	defer outputFile.Close()

	// Execute the template with the data and write the result to the output file
	if err := tmpl.Execute(outputFile, data); err != nil {
		log.Fatal("Failed to execute template:", err)
	}

	log.Println("HTML report generated successfully!")
}
