package main

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/tubbo/cipherutils/cli"
	"github.com/tubbo/cipherutils/dictionary"
)

// Decode a message written in Base64
func Decode(input string, verbose bool) {
	decoded, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		panic(err)
	}

	result := string(decoded)
	words := strings.Split(result, " ")
	count := 0

	for _, word := range words {
		count += dictionary.Lookup(word)
	}

	if verbose || count > 0 {
		fmt.Println(result)
	}
}

func main() {
	cli.Start(Decode)
}
