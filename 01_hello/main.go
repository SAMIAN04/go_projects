package main

import (
	"fmt"

	
)

func main() {
	// Read input
	var action string
	fmt.Scanln(&action)

	// TODO: Define your MediaPlayer interface here
	type MediaPlayer interface {
		Play() string
		Pause() string
		Stop() string
	}
	// TODO: Write your code below to handle the action and print the required output
switch action {
case "play": {
	fmt.Println("MediaPlayer interface requires: Play() string")
}
case "pause" :{
		fmt.Println("MediaPlayer interface requires: Pause() string")

}
case "stop" :{
	fmt.Println("MediaPlayer interface requires: Stop() string")

}
}
	
}
	// Print the result
	// Remember to print in format: "MediaPlayer interface requires: [MethodName]() string"

