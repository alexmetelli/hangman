package main

// Letters_guessed: list (of letters), which letters have been
// guessed so far.
// Returns: string (of letters), comprised of letters that
// represents which letters have not yet been guessed.
func getAvailable(lettersGuessed []string) string {

	var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
		"v", "w", "y", "z"}
	available := ""

	for i := 0; i < len(alphabet); i++ {
		count := 0
		for k := 0; k < len(lettersGuessed); k++ {
			if alphabet[i] == lettersGuessed[k] {
				count++
			}
		}
		if count == 0 {
			available = available + alphabet[i] + " "
		}
		count = 0
	}
	return available
}
