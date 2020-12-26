package main

import (
	"io/ioutil"
	"strings"
)

// Returns a list of valid words. Words are strings of lowercase letters.
// Depending on the size of the word list, this function may take a while to finish.
func loadWords() []string {

	words, err := ioutil.ReadFile("words.txt")
	if err != nil {
		return nil
	}
	s := string(words)

	wordsArray := strings.Split(s, " ")
	return wordsArray
}
