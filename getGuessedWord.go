package main

import "strings"

// Secret_word: string, the word the user is guessing.
// Letters_guessed: list (of letters), which letters have
// been guessed so far.
// Returns: string, comprised of letters, underscores (_), and spaces
// that represents which letters in secret_word have been guessed so far.
func getGuessedWord(secretWord string, letterGuessed []string) string {

	l := len(letterGuessed)
	secretWordList := strings.Split(secretWord, "")
	n := len(secretWordList)
	wordToFill := strings.Repeat("_", n)
	arrayToFill := strings.Split(wordToFill, "")

	for i := 0; i < l; i++ {
		for k := 0; k < n; k++ {
			if letterGuessed[i] == secretWordList[k] {
				arrayToFill[k] = secretWordList[k]
			}
		}
	}
	guessedWord := strings.Join(arrayToFill, " ")
	return guessedWord
}
