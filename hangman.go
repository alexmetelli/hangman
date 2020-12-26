package main

import (
	"fmt"
	"strings"
)

// Secret_word: string, the secret word to guess.
// Starts up an interactive game of Hangman.
// At the start of the game, let the user know how many letters the
// secret_word contains and how many guesses s/he starts with.
// The user should start with 6 guesses.
// Before each round display to the user how many guesses
// s/he has left and the letters that the user has not yet guessed.
// Ask the user to supply one guess per round.
// The user should receive feedback immediately after each guess
// about whether their guess appears in the computer's word.
// After each guess display to the user the partially guessed word so far.
func hangman(secretWord string) {

	secretWordList := strings.Split(secretWord, "")
	l := len(secretWordList)
	guesses := 6
	wordToFill := ""
	arrayToFill := strings.Split(wordToFill, "")
	wordIsGuess := wordIsGuessed(secretWord, arrayToFill)

	fmt.Printf("The secret word has %v letters.\n", l)
	fmt.Printf("You start with %v guesses.\n", guesses)
	fmt.Println("Guess a letter... ")
	var guess string

	// Loops until the word is guessed or no more guesses left.
	for guesses := guesses; guesses > 0; guesses-- {
		fmt.Scanln(&guess)
		// // Checks if input is valid.
		letterCheck := checkString(guess)
		if letterCheck != true {
			for {
				fmt.Println("Error. Only letters accepted ")
				fmt.Println("Guess a letter ")
				fmt.Scanln(&guess)
				letterCheck := checkString(guess)
				if letterCheck == true {
					break
				}
			}
		}

		// Checks if guessed letter is in secret word.
		for i := 0; i < l; i++ {
			if guess == secretWordList[i] {
				wordToFill = wordToFill + guess
				guesses++
				break
			}
			//fmt.Println("Wrong guess")

		}
		arrayToFill = strings.Split(wordToFill, "")
		fmt.Println("Your guess so far: ", getGuessedWord(secretWord, arrayToFill))
		fmt.Println("Available letters:\n", getAvailable(arrayToFill))
		wordIsGuess = wordIsGuessed(secretWord, arrayToFill)
		if wordIsGuess == true {
			fmt.Println("You Won!! Congratulations!!")
			break
		} else if guesses == 0 {
			fmt.Println("You Lost!! No more guesses left!!")
			break
		}
		fmt.Printf("You have %v guesses left\n", guesses)
		fmt.Println()
		fmt.Println("Guess a letter ")
	}

}

// Check if user input is alpha character.
// s: the user input from keyboard.
func checkString(s string) bool {
	if s >= "a" && s <= "z" ||
		s >= "A" && s <= "Z" {
		return true
	}
	return false
}
