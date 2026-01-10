package main

import (
	"fmt"
	"sort"
)

// addWord adds a word to the counter map
func addWord(counter map[string]int, word string) {

	counter[word]++
}

// getCount returns the count for a specific word
func getCount(counter map[string]int, word string) int {
	count := counter[word]
	return count
}

// removeWord removes a word from the counter
// Returns true if the word was found and removed
func removeWord(counter map[string]int, word string) bool {
	_ , exist := counter[word]
	if exist {
		delete(counter,word)
		return exist
	}else{
		return exist
	}

}

// printWordCounts prints the word counts in alphabetical order
func printWordCounts(counter map[string]int) {
	// Get all keys (words) from the map
	words := make([]string, 0, len(counter))
	for word := range counter {
		words = append(words, word)
	}
	
	// Sort the words alphabetically
	sort.Strings(words)
	
	// Print each word and its count in sorted order
	for _, word := range words {
		fmt.Printf("%s: %d\n", word, counter[word])
	}
}

func main() {
	// Initialize the word counter map
	wordCounter := make(map[string]int)
	
	// Add some words
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	for _, word := range words {
		addWord(wordCounter, word)
	}
	
	// Print the word counts
	fmt.Println("Word counts:")
	printWordCounts(wordCounter)
	
	// Get count for specific words
	fmt.Printf("\nCount for 'apple': %d\n", getCount(wordCounter, "apple"))
	fmt.Printf("Count for 'grape': %d\n", getCount(wordCounter, "grape"))
	
	// Remove a word
	removed := removeWord(wordCounter, "banana")
	fmt.Printf("\nRemoved 'banana': %v\n", removed)
	
	// Try to remove a non-existent word
	removed = removeWord(wordCounter, "grape")
	fmt.Printf("Removed 'grape': %v\n", removed)
	
	// Print the final word counts
	fmt.Println("\nFinal word counts:")
	printWordCounts(wordCounter)
}