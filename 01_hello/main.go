package main

import (
	"fmt"
	
	"strconv"
	"strings"
)

type Speaker interface{ Speak() string }
type Person struct{ Name string }
type Parrot struct{ Name string }

func main() {
	// Read input
	var numSpeakersStr string
	var speakerTypesStr string
	var speakerNamesStr string

	fmt.Scanln(&numSpeakersStr)
	fmt.Scanln(&speakerTypesStr)
	fmt.Scanln(&speakerNamesStr)

	// Parse input
	numSpeakers, _ := strconv.Atoi(numSpeakersStr)
	speakerTypes := strings.Split(speakerTypesStr, ",")
	speakerNames := strings.Split(speakerNamesStr, ",")
speakers := make([]Speaker,0)
	// TODO: Write your code below
	// 1. Define the Speaker interface
	// 2. Define Person and Parrot structs
	// 3. Implement Speak() methods for both structs
	// 4. Create makeAllSpeak function
	// 5. Create speakers based on input and store in a slice
	// 6. Call makeAllSpeak with your slice
	for i := 0; i < numSpeakers; i++ {
	    if speakerTypes[i] == "person"{
			speakers = append(speakers, Person{Name: speakerNames[i]})
		}else if speakerTypes[i] =="parrot" {
			speakers = append(speakers, Parrot{Name: speakerNames[i]})
		}else{
			fmt.Println("unknown type")
		}
	}
	makeAllSpeak(speakers)
}
func makeAllSpeak(speakers []Speaker) {
   for _,speak := range speakers {
      fmt.Println(speak.Speak())
   }
}
func (p Person) Speak() string {
  return  fmt.Sprintf("Hello, I'm %s", p.Name)
}
func (p Parrot) Speak() string {
  return fmt.Sprintf("Squawk! %s says hello!", p.Name)
}
