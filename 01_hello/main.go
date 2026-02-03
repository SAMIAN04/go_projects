package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Task struct {
	Name      string
	Completed bool
}

func main()  {
	var numTasksStr string
	var taskData string
	// var completedTaskStr string
	var removeTaskstr string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numTasksStr = scanner.Text()
	scanner.Scan()
	taskData = scanner.Text()
	// scanner.Scan()
	// completedTaskStr = scanner.Text()
	scanner.Scan()
	removeTaskstr = scanner.Text()
	numTasks, _ := strconv.Atoi(numTasksStr)
	// completedTask , _ := strconv.Atoi(completedTaskStr)
	removeTask , _ := strconv.Atoi(removeTaskstr)
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
    // completeTask(&tasks, completedTask) 
	removetask(&tasks,removeTask)
    viewAllTasks(tasks)
	completedCount := 0
	for _, task := range tasks {
		if task.Completed {
			completedCount++
		}
	}
	
	incompleteCount := len(tasks) - completedCount
	fmt.Printf("Task '%v' removed successfully!\n", tasks[removeTask].Name)
	// fmt.Printf("Task '%s' marked as completed!\n", tasks[completedTask].Name)
	fmt.Printf("Total: %d tasks (%d completed, %d remaining)\n", numTasks, completedCount, incompleteCount)
}

func completeTask(tasks *[]Task, index int)  {
	(*tasks)[index].Completed = true
	
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
func removetask(tasks *[]Task, index int) []Task  {
	return append((*tasks)[:index], (*tasks)[index+1:]...)
}