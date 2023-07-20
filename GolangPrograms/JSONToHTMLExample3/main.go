package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>People Data</title>
</head>
<body>
	<table>
		<tr>
			<th>Name</th>
			<th>Age</th>
			<th>Email</th>
		</tr>
		{{range .}}
		<tr>
			<td>{{.Name}}</td>
			<td>{{.Age}}</td>
			<td>{{.Email}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>
`

func main() {
	jsonData := `[{"name":"John Doe","age":30,"email":"john.doe@example.com"},{"name":"Jane Smith","age":25,"email":"jane.smith@example.com"}]`

	var people []Person
	err := json.Unmarshal([]byte(jsonData), &people)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Now the JSON data is parsed into the 'people' slice of 'Person' structs.
	// Next, we'll create an HTML template to display this data.

	tmpl, err := template.New("html").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Error creating template:", err)
		return
	}

	// Finally, we'll render the template with the data and write the result to an output file.
	file, err := os.Create("output.html")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	err = tmpl.Execute(file, people)
	if err != nil {
		fmt.Println("Error rendering template:", err)
		return
	}

	fmt.Println("HTML file 'output.html' generated successfully.")
}
