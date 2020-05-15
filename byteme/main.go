package main

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync"

	"github.com/tubbo/cipherutils/cli"
	"github.com/tubbo/cipherutils/dictionary"
)

var wg sync.WaitGroup

func hexadecimal(input string, verbose bool) {
	output, err := hex.DecodeString(input)

	if err != nil {
		panic(err)
	}

	results := string(output)
	words := strings.Split(results, " ")
	count := 0

	for _, word := range words {
		count += dictionary.Lookup(word)
	}

	if verbose || count > 0 {
		fmt.Println(results)
	}
}

// disabled for now since it's not as simple as hex
func binary(input string, verbose bool) {
	var results string

	for _, c := range input {
		results = fmt.Sprintf("%s%b", results, c)
	}
	words := strings.Split(results, " ")
	count := 0

	for _, word := range words {
		count += dictionary.Lookup(word)
	}

	if verbose || count > 0 {
		fmt.Println(results)
	}
}

// Decode a string from binary to text
func Decode(input string, verbose bool) {
	tokens := strings.Split(input, " ")

	for _, token := range tokens {
		hexadecimal(token, verbose)
		binary(token, verbose)
	}
}

func main() {
	cli.Start(Decode)
}
