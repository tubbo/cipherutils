package main

import (
	"fmt"
	"github.com/markbates/pkger"
	"github.com/yuya-takeyama/argf"
	"io/ioutil"
	"log"
	"strings"
	"sync"
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

func lookup(word string) int {
	file, err := pkger.Open("/dictionary.txt")
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
			fmt.Println(word)
			count++
		}
	}

	return count
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
		found = lookup(downcased)
	}

	if found > 0 {
		fmt.Printf("ROT%d:\t%s", rotation, string(solution))
	}
}

func main() {
	reader, err := argf.Argf()

	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	original := strings.ToUpper(string(buf))

	for rotation := 1; rotation <= 25; rotation++ {
		wg.Add(1)
		go solve(original, rotation)
	}

	wg.Wait()
}
