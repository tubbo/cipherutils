package main

import (
	"encoding/hex"
	"fmt"
	"github.com/tubbo/cipherutils/dictionary"
	"github.com/yuya-takeyama/argf"
	"io/ioutil"
	"strings"
)

func hexadecimal(input string) {
	var output []byte

	bytes := []byte(input)

	_, err := hex.Decode(output, bytes)

	if err != nil {
		panic(err)
	}

	results := string(output)
	words := strings.Split(results, " ")
	count := 0

	for _, word := range words {
		count += dictionary.Lookup(word)
	}

	if count > 0 {
		fmt.Println(results)
	}
}

// disabled for now since it's not as simple as hex
func binary(input string) {
	// results := string(input)
	// words := strings.Split(results, " ")
	// count := 0

	// for _, word := range words {
	// 	count += dictionary.Lookup(word)
	// }

	// if count > 0 {
	// 	fmt.Println(results)
	// }
}

// Decode a string from binary to text
func Decode(input string) {
	go hexadecimal(input)
	go binary(input)
}

func main() {
	reader, err := argf.Argf()
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	input := string(buf)

	Decode(input)
}
