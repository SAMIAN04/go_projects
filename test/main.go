package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Opatrations struct {
	Task  string
	Type  string
	Name  string
	score int
}

func main() {
	// Read input
	var numOpsStr string
	var operationsStr string
	fmt.Scanln(&numOpsStr)
	fmt.Scanln(&operationsStr)
	//numOps, _ := strconv.Atoi(numOpsStr)
	// Declare maps without initialization (they will be nil)
	var scores map[string]int
	var grades map[string]int
	fmt.Println(scores, grades)
	// Split operations by comma
	operationsParts := strings.Split(operationsStr, ",")
	operations := []Opatrations{}

	
	for _, operationParts := range operationsParts {
		operation := strings.Split(operationParts, ":")

		if len(operation) == 2 {
			task := operation[0]
			types := operation[1]
					    operations = append(operations,Opatrations{Task: task,Type: types,} )

		} else if len(operation) == 4 {
			task := operation[0]
			types := operation[1]
			name := operation[2]
		    score ,_:= strconv.Atoi(operation[3])
		    operations = append(operations,Opatrations{Task: task,Type: types,score: score,Name: name} )
		} 
		
		//    task := operation[0]
		//    types := operation[1]
		//    name := operation[2]
		//    score ,_:= strconv.Atoi(operation[3])
		//    operations = append(operations,Opatrations{Task: task,Type: types,score: score,Name: name} )

	}
	fmt.Println(operations)

}

func Mapinit(scores map[string]int, oparetions []Opatrations)  {
	if scores  {
		
	}
}