package cli

// Package cli wraps the command-line interface of cipherutils programs
// into a reusable component that can be given a callback function to
// execute once input from the user is detected. This enables all
// cipherutils programs to use the same command-line interface of
// reading a file path or STDIN.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yuya-takeyama/argf"
)

var verbose = flag.Bool("v", false, "Verbose mode, show all solutions")

// Start a new CLI
func Start(callback func(input string, verbose bool)) {
	flag.Usage = func() {
		name := os.Args[0]
		output := flag.CommandLine.Output()

		fmt.Fprintf(output, "%s, part of the cipherutils - https://tubbo.github.io/cipherutils\n\n", name)
		fmt.Fprintf(output, "Usage: %s [-hv] ./path/to/file\n", name)
		fmt.Fprintf(output, "Flags:\n")
		fmt.Fprintf(output, "  -h	Show this help\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	reader, err := argf.From(flag.Args())

	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	input := string(buf)

	callback(input, *verbose)
}
