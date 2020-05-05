package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/tubbo/cipherutils/cli"
	"github.com/tubbo/cipherutils/dictionary"
)

var wg sync.WaitGroup
var alphabet = [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func index(value rune) int {
	for index, character := range alphabet {
		if character == value {
			return index
		}
	}
	return -1
}

func rotate(position int, rotation int) int {
	rotated := position + rotation

	if rotated > len(alphabet) {
		rotated -= len(alphabet)
	}

	return rotated - 1
}

func solve(text string, rotation int) {
	defer wg.Done()
	var solution []rune

	for _, character := range text {
		position := index(character)

		if position == -1 {
			solution = append(solution, character)
		} else {
			rotated := rotate(position, rotation)
			solution = append(solution, alphabet[rotated])
		}
	}

	words := strings.Split(string(solution), " ")
	found := 0

	for _, word := range words {
		trimmed := strings.TrimSuffix(word, "\n")
		downcased := strings.ToLower(trimmed)
		found = dictionary.Lookup(downcased)
	}

	if found > 0 {
		fmt.Printf("ROT%d:\t%s", rotation, string(solution))
	}
}

// Decode solves a rotational cipher by running the text through all 25
// permutations, and outputting the result that matches words in the
// dictionary.
func Decode(input string) {
	encoded := strings.ToUpper(input)

	for rotation := 1; rotation <= 25; rotation++ {
		wg.Add(1)
		go solve(encoded, rotation)
	}

	wg.Wait()
}

func main() {
	cli.Start(func(input string) {
		Decode(input)
	})
}
