package main

import "strings"

// Secret_word: string, the word the user is guessing;
// assumes all letters are lowercase.
// Letters_guessed: list (of letters), which letters have been
// guessed so far; assumes that all letters are lowercase.
// Returns: boolean, True if all the letters of secret_word are
// in letters_guessed; False otherwise.
func wordIsGuessed(secretWord string, lettersGuessed []string) bool {

	guessedWord := getGuessedWord(secretWord, lettersGuessed)
	wordStripped := strings.ReplaceAll(guessedWord, " ", "")
	return secretWord == wordStripped
}
