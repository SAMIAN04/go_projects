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
	var languagesStr string
	fmt.Scanln(&countStr)
	fmt.Scanln(&languagesStr)
	// Convert count string to integer (not needed for this challenge)
	_, _ = strconv.Atoi(countStr)

	// Split the languages string by commas
	languages := strings.Split(languagesStr, ",")

	// Create a set using map[string]struct{} idiom
	languageSet := make(map[string]struct{})

	for _, language := range languages {

		_, exist := languageSet[language]
		languageSet[language] = struct{}{}
		if exist {
			fmt.Printf("Already exists: %s\n", language)
		} else {
			fmt.Printf("Added: %s\n", language)
		}

	}
	fmt.Printf("Total languages processed: %s\n", countStr)
	fmt.Printf("Unique languages: %d\n", len(languageSet))

	// Print all unique languages in the set
	fmt.Println("Programming languages in set:")
	keys := make([]string, 0, len(languageSet))
	for k := range languageSet {
		keys = append(keys, k)

	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("- %s\n", k)
	}
	// Convert map to slice and sort for consistent output
	// List each language with "- [language]" format
}
