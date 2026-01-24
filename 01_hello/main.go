package main

import (
	"fmt"
	
	"strconv"
	"strings"
)

func main() {
	// Read input
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)

	// TODO: Write your code below
	// 1. Create the processData function that accepts interface{}
	// 2. Convert valueStr to the appropriate type based on dataType
	// 3. Call processData with the converted value

	switch dataType {
	case "int":
		{
			intiger, _ := strconv.Atoi(valueStr)
			processData(intiger)
		}
	case "string":
		{
			processData(valueStr)
		}
	case "bool": {
		b, _ := strconv.ParseBool(valueStr)
		processData(b)
	}
case "slice" : {
	
	parts := strings.Split(valueStr, ",")
	slices := make([]int, 0,len(parts))
	for _, part := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil {
			slices = append(slices, n)
		}
	}
	processData(slices)
	
}
	}
}
func processData(v interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", v,v)
}
