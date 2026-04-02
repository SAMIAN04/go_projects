package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var countStr string
	var skillsStr string
	fmt.Scanln(&countStr)
	fmt.Scanln(&skillsStr)
	
	// Convert count string to integer (for reference only - not used in loop)
	_, _ = strconv.Atoi(countStr)
	
	// Initialize skill set with starter skills
	skillSet := map[string]struct{}{
		"Programming":     struct{}{},
		"Problem Solving": struct{}{},
		"Communication":   struct{}{},
	}
	
	// Split skills string into individual skills
	skills := strings.Split(skillsStr, ",")
	
	// TODO: Write your code below
	// Process each skill using range to iterate over ALL skills in the slice
	// Use comma ok idiom: _, exists := skillSet[skill]
	// Add new skills using: skillSet[skill] = struct{}{}
	// Display processing results for each skill
	// Calculate and display progress summary using len(skills)
	// List all skills in the complete skill set (consider sorting for consistent output)
	
}