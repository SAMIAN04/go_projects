package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Name      string
	completed bool
}

func main() {
	Tasks := make([]Task, 0)

	var num int
	fmt.Scanln(&num)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := scanner.Text()
	multpleTasks := strings.Split(inputs, ",")

	if inputs != "" {
		for i := 0; i < len(multpleTasks); i++ {
			Tasks = addTask(Tasks, multpleTasks[i])
		}
	}

	// Print all tasks
	//viewAllTasks(Tasks)

	// Count completed and remaining
	completed := 0
	for _, task := range Tasks {
		if task.completed {
			completed++
		}
	}
	remaining := len(Tasks) - completed

	fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", len(Tasks), completed, remaining)
}

// addTask
func addTask(tasks []Task, taskName string) []Task {
	tasktype := strings.Split(taskName, ":")
	
	// Extract just the name (before the colon)
	name := ""
	isCompleted := false
	if len(tasktype) > 0 {
		name = strings.TrimSpace(tasktype[0])
	}
	if len(tasktype) > 1 {
		isCompleted = strings.TrimSpace(tasktype[1]) == "true"
	}
	
	task := Task{Name: name, completed: isCompleted}
	tasks = append(tasks, task)
	
	return tasks
}

func viewAllTasks(Tasks []Task) {
	for _, task := range Tasks {
		if task.completed {
			fmt.Printf("[x] %s\n", task.Name)
		} else {
			fmt.Printf("[ ] %s\n", task.Name)
		}
	}
}
func completeTask(Task , index int) string  {
	return ""
}