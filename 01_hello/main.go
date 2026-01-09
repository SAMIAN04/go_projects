package main

import "fmt"

func main() {
	// A map of student grades
	studentGrades := map[string]string{
		"John":  "B+",
		"Emma":  "A-",
		"Sarah": "A",
		"Mike":  "C",
	}
	
	// TODO: Access and store Emma's grade in a variable called emmaGrade
	emmaGrade := studentGrades["Emma"]
	// TODO: Print Emma's grade with a message like: "Emma's grade is: A-"
	fmt.Printf("Emma's grade is: %s\n",emmaGrade)
}