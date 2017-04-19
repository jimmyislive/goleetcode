// https://leetcode.com/problems/detect-capital/#/description
/*
Given a word, you need to judge whether the usage of capitals in it is right or not.

We define the usage of capitals in a word to be right when one of the following cases holds:

All letters in this word are capitals, like "USA".
All letters in this word are not capitals, like "leetcode".
Only the first letter in this word is capital if it has more than one letter, like "Google".
Otherwise, we define that this word doesn't use capitals in a right way.

*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func detectCapitalUse(word string) bool {

	if word == strings.ToUpper(word) {
		return true
	}

	if word == strings.ToLower(word) {
		return true
	}

	if len(word) > 1 && string(word[0]) == strings.ToUpper(string(word[0])) &&
		string(word[1:]) == strings.ToLower(string(word[1:])) {
		return true
	}

	return false
}

func main() {
	fmt.Println(detectCapitalUse(os.Args[1]))
}
