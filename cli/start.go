package cli

// Package cli wraps the command-line interface of cipherutils programs
// into a reusable component that can be given a callback function to
// execute once input from the user is detected. This enables all
// cipherutils programs to use the same command-line interface of
// reading a file path or STDIN.

import (
	"github.com/yuya-takeyama/argf"
	"io/ioutil"
	"log"
)

// Start a new CLI
func Start(callback func(input string)) {
	reader, err := argf.Argf()

	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	input := string(buf)

	callback(input)
}
