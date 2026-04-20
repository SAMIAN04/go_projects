package main

import (
	"fmt"
	"sort"
	//"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	var countStr string
	var skillsStr string
	var allSkills []string
	fmt.Scanln(&countStr)
	fmt.Scanln(&skillsStr)
	var skillMasterd int
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
	for _,skill := range skills{
		    
			
			
			if _,exist := skillSet[skill]; exist  {
				fmt.Printf("Already mastered: %s\n", skill)
				skillMasterd+=1
			}else {
				fmt.Printf("Learning new skill: %s\n", skill)
				skillSet[skill] = struct{}{}
			}

	}
	
	fmt.Printf("Skills processed: %d\n", len(skills) )
	fmt.Printf("New skills learned: %d\n", len(skillSet)-3)
	fmt.Printf("Total skills mastered: %d\n", len(skillSet))
	fmt.Println("Complete skill set:")
	for skill := range skillSet{
		allSkills = append(allSkills, skill)
	}
sort.Strings(allSkills)
for _,skill := range allSkills {
	fmt.Printf("✓ %s\n", skill)
}


}