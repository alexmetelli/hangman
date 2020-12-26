package main

import (
	"math/rand"
	"time"
)

func chooseWord(wordList []string) string {

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(wordList)
	return wordList[n]
}
