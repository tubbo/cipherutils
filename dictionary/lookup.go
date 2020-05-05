package dictionary

// Package dictionary implements a simple dictionary that one can look
// up words in. It's used in cipherutils to find the correct solution
// for a given cipher.

import (
	"github.com/markbates/pkger"
	"io/ioutil"
	"strings"
)

// Lookup a word in the dictionary and return how many times it appears
func Lookup(word string) int {
	file, err := pkger.Open("words.txt")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	words := strings.Split(string(data), "\n")
	count := 0

	if err != nil {
		panic(err)
	}

	for _, dictword := range words {
		if dictword == word {
			count++
		}
	}

	return count
}
