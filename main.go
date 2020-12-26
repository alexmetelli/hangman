package main

import (
	"fmt"
	"strings"
)

func main() {

	// Print titles.
	welcome := "*  WELCOME TO THE HANGMAN...  *"
	star := strings.Repeat("*", len(welcome))
	fmt.Println()
	fmt.Println(star)
	fmt.Println()
	fmt.Println(welcome)
	fmt.Println()
	fmt.Println(star)
	fmt.Println()

	words := loadWords()
	secretWord := chooseWord(words)

	//letters := []string{"o", "k", "l", "d", "w", "n"}

	//loaded := loadWords()
	//word := chooseWord(loaded)
	//guessed := getGuessedWord(secretWord, letters)
	//available := getAvailable(letters)
	//wordGuessed := wordIsGuessed(secretWord, letters)
	// fmt.Printf("%t", wordGuessed)
	hangman(secretWord)
}
