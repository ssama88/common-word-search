package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	strings := [10]string{"Here I am", "There I am am am am am am am am am", "What I am"}
	getCommonWordsInStrings(strings)
}

func getCommonWordsInStrings(strings [10]string) {
	if !validateInput(strings) {
		return
	}
	uniqueWords := make(map[string]int)

	for i := range strings {
		if len(strings[i]) == 0 { //check for potential empty strings in input
			continue
		}
		processText(strings[i], uniqueWords)
	}

	for key, val := range uniqueWords {
		fmt.Printf("%s: %d\n", key, val)
	}

}

// Can't compare more than 10 strings
func validateInput(strings [10]string) bool {
	if len(strings) > 10 {
		return false
	} else if len(strings) == 1 {
		return false
	}
	return checkValidStringLength(strings)
}

// string can't be more than 1000 characters
func checkValidStringLength(strings [10]string) bool {
	for i := range strings {
		if len(strings[i]) > 1000 {
			return false
		}
	}
	return true
}

func processText(text string, uniqueWords map[string]int) {

	localUniqueWords := make(map[string]bool)

	var words []string = strings.Split(text, " ")

	for i := range words {
		_, inLocalText := localUniqueWords[words[i]]
		if inLocalText {
			continue
		} else {
			localUniqueWords[words[i]] = true
		}
		_, inGlobalWords := uniqueWords[words[i]]
		if inGlobalWords {
			uniqueWords[words[i]]++
		} else {
			uniqueWords[words[i]] = 1
		}
	}

}
