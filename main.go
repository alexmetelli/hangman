package main

func main() {

	secretWord := "lockdown"
	//letters := []string{"o", "k", "l", "d", "w", "n"}

	//loaded := loadWords()
	//word := chooseWord(loaded)
	//guessed := getGuessedWord(secretWord, letters)
	//available := getAvailable(letters)
	//wordGuessed := wordIsGuessed(secretWord, letters)
	// fmt.Printf("%t", wordGuessed)
	hangman(secretWord)
}
