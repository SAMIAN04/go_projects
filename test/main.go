package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Studentinfo struct {
	Id    int
	Grade string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	studentNumberStr := scanner.Text()
	scanner.Scan()
	studentsDataStr := scanner.Text()
	studentNumber, _ := strconv.Atoi(studentNumberStr)
	//  gradesExist := false
	studentData := make(map[string]Studentinfo)
	gradesCount := make(map[string]int)
	studentsData := strings.Split(studentsDataStr, ",")
	maxName := ""
	maxId := 0
	//grades := []Studentinfo{}
	keys := []string{}
     grades := []string{}
	for i := 0; i < studentNumber; i++ {
		parts := strings.Split(studentsData[i], ":")
		Name := parts[0]
		Id, _ := strconv.Atoi(parts[1])
		Grade := parts[2]
		studentData[Name] = Studentinfo{
			Id:    Id,
			Grade: Grade,
		}

	}
	

	for k := range studentData {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for i := range keys {

		fmt.Printf("%s: ID %d, Grade %s\n", keys[i], studentData[keys[i]].Id, studentData[keys[i]].Grade)
		gradesCount[studentData[keys[i]].Grade]++
		if studentData[keys[i]].Id > maxId {
			maxId = studentData[keys[i]].Id
			maxName = keys[i]
		}
	}
	// we dk how tf it works we neeed to find out amd lot of works left

	for k := range gradesCount {
		
		grades =append(grades,k)
		
	}
	sort.Slice(grades,func(i, j int) bool {
		return grades[i]< grades[j]
	})
	for _,g := range grades {
           fmt.Printf("Grade %s: %d students\n", g,gradesCount[g])
	}
	
	fmt.Printf("Highest ID: %s (%d)\n", maxName,maxId)
    fmt.Printf("Total students: %d\n", studentNumber)
}
