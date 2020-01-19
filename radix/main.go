package main

import (
	"encoding/base64"
	"fmt"
	"github.com/tubbo/cipherutils/dictionary"
	"github.com/yuya-takeyama/argf"
	"io/ioutil"
	"strings"
)

// Decode a message written in Base64
func Decode(input string) {
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

	if count > 0 {
		fmt.Println(result)
	}
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
