package main

import (
	"bufio"
	"fmt"
	"strings"

	"os"
)
type Task struct {
		Name      string
		completed bool
	}
func main() {
	// Read input
	var appName string
	var userName string
	fmt.Scanln(&appName)
	fmt.Scanln(&userName)
	var num int
// TODO: Write your code below
	// Define the Task struct
	// Create a slice of Task structs
	// Display welcome message and menu
	// Print current tasks count
	fmt.Printf("Welcome to %s, %s!\n", appName, userName)
	Tasks := make([]Task, 0)
	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Remove Task")
	fmt.Println("5. Exit")
	fmt.Printf("Current tasks: %d\n",len(Tasks) )
    fmt.Scanln(&num)
// if num == 1 {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	inputs := scanner.Text()
// 	multpleTasks := strings.Split(inputs, ",")
// 	for i := 0; i < len(multpleTasks); i++ {
// 	Tasks =	addTask(Tasks,multpleTasks[i])
// 	}
	
// 	}else {
// 		fmt.Println("enter valid option")
// }

fmt.Printf("Current tasks: %d\n",len(Tasks) )
	

}

//addTask 
func addTask(tasks []Task, taskName string) []Task {
     task := Task{Name: taskName, completed: false}
	 tasks = append(tasks, task)
	 return tasks
}