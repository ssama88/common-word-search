package main

import (
	"fmt"
	"strings"
)

func main() {
	getMatchingStrings("Hello There World", "World")
	getMatchingStrings("Snake my in", "Theres a snake in my boot")
}

// I want to find the common words in 2 different walls of text

//It would be more effecient to go through the shorter string because that's the maximum amount of words that can be found in both strings

/*EX: The star saw the moon
  saw the moon

  "The star" is irrelevant and would need 5 computations to map through while "saw the moo would only need 3"
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
