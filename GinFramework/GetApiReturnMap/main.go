package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	router := gin.Default()

	router.GET("/data", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":  "John Doe",
			"age":   30,
			"email": "john.doe@example.com",
		}

		c.JSON(http.StatusOK, data)

		fmt.Println(data)

		jsonByte, err := json.Marshal(data)
		if err != nil {
			fmt.Println("marshal error ", err)
		}

		// Parse JSON data into the Data struct
		var persondata Person
		if err := json.Unmarshal([]byte(jsonByte), &persondata); err != nil {
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

	})

	router.Run(":8083")
}
