package main

import "fmt"

func main() {
	// Person struct with three fields
	type Person struct {
		name      string
		age       int
		isStudent bool
	}

	// A sample person
	john := Person{
		name:      "John Doe",
		age:       25,
		isStudent: true,
	}

	// TODO: Print john's name using fmt.Printf and the %s format verb
	fmt.Printf("Name: %s\n", john.name)
	
	// TODO: Print john's age using fmt.Printf and the %d format verb
	fmt.Printf("Age: %d\n", john.age)
	
	// TODO: Print whether john is a student using fmt.Printf and the %t format verb
	fmt.Printf("Is student: %t\n", john.isStudent)
}