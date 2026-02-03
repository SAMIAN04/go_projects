package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func viewAllTasks(tasks []Task) {
	for _, task := range tasks {
		if task.Completed {
			fmt.Printf("[x] %s\n", task.Name)
		} else {
			fmt.Printf("[ ] %s\n", task.Name)
		}
	}
}

func main() {
	var numTasksStr string
	var taskData string
	fmt.Scan(&numTasksStr)
	fmt.Scan(&taskData)
	
	numTasks, _ := strconv.Atoi(numTasksStr)
	
	var tasks []Task
	taskEntries := strings.Split(taskData, ",")
	
	for _, entry := range taskEntries {
		parts := strings.Split(entry, ":")
		if len(parts) >= 2 { // Add safety check
			name := parts[0]
			completed, _ := strconv.ParseBool(parts[1])
			
			task := Task{
				Name:      name,
				Completed: completed,
			}
			tasks = append(tasks, task)
		}
	}
	
	viewAllTasks(tasks)
	
	completedCount := 0
	for _, task := range tasks {
		if task.Completed {
			completedCount++
		}
	}
	
	incompleteCount := numTasks - completedCount
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", numTasks, completedCount, incompleteCount)
}