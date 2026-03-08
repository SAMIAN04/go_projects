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

	// Split operations by comma
	operationsParts := strings.Split(operationsStr, ",")
	operations := []Opatrations{}

	for _, operationParts := range operationsParts {
		operation := strings.Split(operationParts, ":")

		if len(operation) == 2 {
			task := operation[0]
			types := operation[1]
			operations = append(operations, Opatrations{Task: task, Type: types})

		} else if len(operation) == 4 {
			task := operation[0]
			types := operation[1]
			name := operation[2]
			score, _ := strconv.Atoi(operation[3])
			operations = append(operations, Opatrations{Task: task, Type: types, score: score, Name: name})
		}
	}
	for _, operation := range operations {
		switch operation.Task {
		case "init":
			{
				if operation.Type == "scores" {
					Mapinit(scores, operation.Type)
				} else if operation.Type == "grades" {
					Mapinit(grades, operation.Type)

				}

			}
		case "add" :{
			if operation.Type == "scores" {
					add(scores,operation.Type,operation.Name,operation.score)
				} else if operation.Type == "grades" {
					add(grades,operation.Type,operation.Name,operation.score)

				}
		}
		}
	}

}

func Mapinit(scores map[string]int, Type string) {
	if scores == nil {
		scores = make(map[string]int)
	}
	fmt.Printf("Initialized map %s\n", Type)
}

func add (scores map[string]int,Type string, name string,score int ){
   if scores == nil {
	fmt.Printf("Cannot add to nil map %s\n", Type)
   }else {
	fmt.Printf("Added %s:%d to %s\n", name,score,Type)
	scores[name] = score
	
   }
}
