package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/yuya-takeyama/argf"
	"io/ioutil"
)

var encoding = flag.Bool("h", false, "Decode from hexadecimal")

func hexadecimal(input string) string {
	var output []byte

	bytes := []byte(input)

	hex.Decode(output, bytes)

	return string(output)
}

func binary(input string) string {
	return input
}

// Decode a string from binary to text
func Decode(input string, isHex bool) string {
	if isHex {
		return hexadecimal(input)
	}

	return binary(input)
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
	result := Decode(input, true)

	fmt.Println(result)
}
