package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read the full line of text (including spaces)
	scanner.Scan()
	text := scanner.Text()

	// Read the threshold value
	scanner.Scan()
	thresholdStr := scanner.Text()

	// Convert threshold to integer
	threshold, _ := strconv.Atoi(thresholdStr)
	wordMap := make(map[string]int)
	UniquWordsSlice := []string{}
	AbovThreshold := []string{}

	var mostFrequentWordCount int
	mostFrequentWord := ""
	fmt.Println("Word Frequency Analysis:")
	countWords(text, threshold, wordMap, &UniquWordsSlice, &AbovThreshold, &mostFrequentWordCount, &mostFrequentWord)
	fmt.Printf("Total unique words: %d\n", len(wordMap))
	fmt.Printf("Words above threshold: %d\n", len(AbovThreshold))
	fmt.Printf("Most frequent word: %s (%d times)\n", mostFrequentWord, mostFrequentWordCount)
}

func countWords(text string, threshold int, wordMap map[string]int, UniquWordsSlice *[]string, AbovThreshold *[]string, mostFrequentWordCount *int, mostFrequentWord *string) {
	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		wordMap[word]++
	}

	for word, count := range wordMap {
		// fmt.Printf("%s: %d\n", word, count)
		if count == 1 {
			*UniquWordsSlice = append(*UniquWordsSlice, word)
		}
		if count >= threshold {
			*AbovThreshold = append(*AbovThreshold, word)
		}
		
	}

	sort.Strings(*AbovThreshold)
	for _, word := range *AbovThreshold {
		fmt.Printf("%s: %d\n", word, wordMap[word])
		if wordMap[word] > *mostFrequentWordCount && wordMap[word] != *mostFrequentWordCount {
			*mostFrequentWordCount = wordMap[word]
			*mostFrequentWord = word

		}
	}
}
