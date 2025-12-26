package main

import "fmt"

func main() {
	// Declare a variable to store the user's name
	var name string
	
	// Display a prompt for the user
	fmt.Print("Enter your name: ")
	
	// TODO: Use fmt.Scanln() to read the user's input into the 'name' variable
	fmt.Scanln(&name)
	// Display a greeting using the name entered
	fmt.Printf("Hello, %s! fuck me harder.\n", name)
}