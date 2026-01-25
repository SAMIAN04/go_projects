package main

import (
	"fmt"
	
	"strconv"
)

func main() {
	// Read input
	var dataType string
	var valueStr string
	fmt.Scanln(&dataType)
	fmt.Scanln(&valueStr)

	// Variable to store the interface value
	var interfaceValue interface{}
	switch dataType {
	case "int":
		{
			interfaceValue, _ = strconv.Atoi(valueStr)
			value, ok := interfaceValue.(int)
	if ok {
		 fmt.Printf("Success: %v is a %T\n", value, value)
	} else {
		fmt.Printf("Failed: value is not a %s\n", dataType)
	}


		}
	case "string":
		{
			interfaceValue = valueStr
		value, ok := interfaceValue.(string)
	if ok {
		 fmt.Printf("Success: %v is a %T\n", value, value)
	} else {
		fmt.Printf("Failed: value is not a %s\n", dataType)
	}
		}
	case "bool":
		{
			interfaceValue, _ = strconv.ParseBool(valueStr)
		value, ok := interfaceValue.(bool)
	if ok {
		 fmt.Printf("Success: %v is a %T\n", value, value)
	} else {
		fmt.Printf("Failed: value is not a %s\n", dataType)
	}
		}
	}
	// TODO: Write your code below
	// 1. Convert valueStr to appropriate type based on dataType and store in interfaceValue
	// 2. Use type assertion to check if interfaceValue contains the expected type
	// 3. Print the appropriate success or failure message

}