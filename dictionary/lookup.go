package dictionary

// Package dictionary implements a simple dictionary that one can look
// up words in. It's used in cipherutils to find the correct solution
// for a given cipher.

// Lookup a word in the dictionary and return how many times it appears
func Lookup(word string) int {
	count := 0

	for _, dictword := range Words {
		if dictword == word {
			count++
		}
	}

	return count
}
