package main

import (
	"fmt"
	"strings"
)

func main() {
	strings := [10]string{"Here I am", "There I am am am am am am am am am", "What I am"}
	getCommonWordsInStrings(strings)
}

// I want to find the common words in 2 different walls of text

//It would be more effecient to go through the shorter string because that's the maximum amount of words that can be found in both strings

/*EX: The star saw the moon
  saw the moon

  "The star" is irrelevant and would need 5 computations to map through while "saw the moo would only need 3"

  For more than one string maybe look into https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_string-search_algorithm
*/

func getMatchingStrings(string1 string, string2 string) {
	var smallerString string
	var biggerString string
	if len(string1) > len(string2) {
		smallerString = string2
		biggerString = string1
	} else {
		smallerString = string1
		biggerString = string2
	}

	var words []string = strings.Split(smallerString, " ")
	var words2 []string = strings.Split(biggerString, " ")

	var commonWords []string = make([]string, 0)
	for i := range words {
		for j := range words2 {
			if strings.EqualFold(words[i], words2[j]) {
				commonWords = append(commonWords, words[i])
				break //only want common words, dont care about instances right now
			}
		}
	}

	for i := range commonWords {
		fmt.Printf("%s \n", strings.ToLower(commonWords[i]))
	}
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
	//Here : 1, I: 4 , am: 3 , There:1, What:1
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
		if words[i] == "" {
			fmt.Println(("why?"))
		}
		_, inLocalText := localUniqueWords[words[i]]
		if inLocalText {
			continue
		} else {
			localUniqueWords[words[i]] = true
		}
		test, inGlobalWords := uniqueWords[words[i]]
		if inGlobalWords {
			uniqueWords[words[i]]++
		} else {
			fmt.Printf("Key: %s, index: %d, underscore val %d \n", words[i], i, test)
			uniqueWords[words[i]] = 1
		}
	}

}
